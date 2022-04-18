/* 
 *
 * UI functions 
 *
 */

function checkTab(tid) {
   if (tid == "tab-2") {
      getImages();
   }

   if (tid == "tab-4") {
      getEnvs();
   }

   if (tid == "tab-5") {
      getImgDDM();
   }

   document.getElementById(tid).checked = true;
}

async function loadImages() {
   await getImages();
   document.getElementById("tab-2").checked = true;
}

async function editImage(id) {
   await getImage(id);
}

async function loadEnvs() {
   await getEnvs();
   document.getElementById("tab-4").checked = true;
}

async function editEnv(id) {
   await getEnv(id);
}

function sleep(ms) {
   return new Promise(resolve => setTimeout(resolve, ms));
}

function call_popup(text,time,color,background) {

   var html_content = `<div id="popup-container" class="popup"><div style="color:`+color+`;background-color:`+background+`;" class="popup-content"><span class="popup-close">&times;</span>`+text+`</div></div>`;
   document.getElementById("popup-area").innerHTML = html_content;
   var popup = document.getElementById('popup-container');
   var span = document.getElementsByClassName("popup-close")[0];
   popup.style.display = "block";

   span.onclick = function() {
      popup.style.display = "none";
   }
   window.onclick = function(event) {
      if (event.target == popup) {
          popup.style.display = "none";
      }
   }
   if (time != 0) {
     setTimeout(function(){
         popup.style.display = "none";
     }, time);
   }
}

function sortTable(n, t) {
  var table, rows, switching, i, x, y, shouldSwitch, dir, switchcount = 0;
  table = document.getElementById(t);
  switching = true;
  // Set the sorting direction to ascending:
  dir = "asc";
  /* Make a loop that will continue until
  no switching has been done: */
  while (switching) {
    // Start by saying: no switching is done:
    switching = false;
    rows = table.rows;
    /* Loop through all table rows (except the
    first, which contains table headers): */
    for (i = 1; i < (rows.length - 1); i++) {
      // Start by saying there should be no switching:
      shouldSwitch = false;
      /* Get the two elements you want to compare,
      one from current row and one from the next: */
      x = rows[i].getElementsByTagName("TD")[n];
      y = rows[i + 1].getElementsByTagName("TD")[n];
      /* Check if the two rows should switch place,
      based on the direction, asc or desc: */
      if (dir == "asc") {
        if (x.innerHTML.toLowerCase() > y.innerHTML.toLowerCase()) {
          // If so, mark as a switch and break the loop:
          shouldSwitch = true;
          break;
        }
      } else if (dir == "desc") {
        if (x.innerHTML.toLowerCase() < y.innerHTML.toLowerCase()) {
          // If so, mark as a switch and break the loop:
          shouldSwitch = true;
          break;
        }
      }
    }
    if (shouldSwitch) {
      /* If a switch has been marked, make the switch
      and mark that a switch has been done: */
      rows[i].parentNode.insertBefore(rows[i + 1], rows[i]);
      switching = true;
      // Each time a switch is done, increase this count by 1:
      switchcount ++;
    } else {
      /* If no switching has been done AND the direction is "asc",
      set the direction to "desc" and run the while loop again. */
      if (switchcount == 0 && dir == "asc") {
        dir = "desc";
        switching = true;
      }
    }
  }
}

/*
 *
 * Backend
 *
 */

// Test function
window.test = function() {
   var data = {};

   try {
      window.go.main.App.Test()
         .then(result => {
            data = result;
            console.log(data);
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
};

/*
 * IMAGES
 */

// Add image function
window.addImage = function(name, repo) {
   var data = {};

   if (name ==  "" || repo == "") {
      document.getElementById("add-img-form").reset();
      return
   }

   try {
      window.go.main.App.AddImage(name, repo)
         .then(result => {
            data = result;
            document.getElementById("tab-2").checked = true;
            getImages();
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
   document.getElementById("add-img-form").reset();
};

// Change image function
window.changeImage = function(id, name, repo) {
   var data = {};

   if (name ==  "" || repo == "") {
      document.getElementById("edit-img-form").reset();
      return
   }

   try {
      window.go.main.App.ChangeImage(id, name, repo)
         .then(result => {
            data = result;
            getImages();
            document.getElementById("tab-2").checked = true;
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
   document.getElementById("edit-img-form").reset();
};

// Remove image function
window.removeImage = function(id) {
   var data = {};

   try {
      window.go.main.App.RemoveImage(id)
         .then(result => {
            data = result;
            getImages();
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
};

window.getImage = function(id) {

   try {
      data = window.go.main.App.GetImage(id)
         .then(result => {
            img = result;
            var id = img.ID;
            var name = img.Name;
            var repo = img.Repo;

            part1 = `<div class="grid grid-2">
                        <input type="text" id="img-edit-name" value="`;
            part2 = `" required>
                        <input type="text" id="img-edit-repo" value="`;
            part3 = `" required>  
                     </div>`;

            var form = part1 + name + part2 + repo + part3;

            document.getElementById("image-form").innerHTML = form;

            btn1 = `<button type="submit" class="btn-grid" 
                               onclick="changeImage('`;

            btn2 = `', document.getElementById('img-edit-name').value, document.getElementById('img-edit-repo').value)">`;

            footer = `     <span class="back">
                             <img src="./assets/images/save-scaled.svg" alt="">
                           </span>
                           <span class="front"> SAVE </span>
                         </button>`

            var button = btn1 + id + btn2 + footer;

            document.getElementById("img-edit-save").innerHTML = button;

            document.getElementById("tab-2.5").checked = true;
         })
         .catch((err) => {
            console.error(err);
         });
      } catch (err) {
         console.error(err);
      }
};

window.getImages = function() {

   try {
      data = window.go.main.App.GetImages()
         .then(result => {
            imgs = result;
            header = `<table id="image-table">
                   <thead>
                     <tr>  
                       <th>image <button onclick="sortTable(0, 'image-table')" class="t-nav"><img src="./assets/images/d-arrow.png"></button></th>
                       <th>repo:tag <button onclick="sortTable(1, 'image-table')" class="t-nav"><img src="./assets/images/d-arrow.png"></button></th>
                       <th class="t-center">edit</th>
                       <th class="t-center">remove</th>
                     </tr>
                   </thead>
                   <tbody>`;
            part1 =`      
                     <tr>
                       <td>`;
            part2 = `  </td>
                        <td>`;
            part3 = `</td>
                       <td class="t-center"><button onclick="editImage('`; 
            part4 = `')" class="table-btn">edit</button></td>
                       <td class="t-center"><button onclick="removeImage('`;
            part5 = `')" id="rm-`;

            part6 = `" class="table-btn remove-btn"><img class="remove-icon" src="./assets/images/remove.png"></button></td>
                     </tr>`;
                     //<tr class="spacer"><td colspan="100"></td></tr>`;
            footer =`
                     </tbody>
                        </table>
                        <button onclick="checkTab('tab-3')" class="add-btn">+</button>`;

            var body = ""

            for (let i = 0; i < imgs.length; i++) {
               id = imgs[i].ID
               name = imgs[i].Name
               repo = imgs[i].Repo
               body = body + part1 + name + part2 + repo + part3 + id + part4 + id + part5 + id + part6;
            }

            var table = header + body + footer;

            document.getElementById("images-table").innerHTML = table;

         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }

};

// Get all images for dropdown menu
window.getImgDDM = function() {
   
   try {
      data = window.go.main.App.GetImages()
         .then(result => {
            imgs = result;

            header = `<select id="env-image" required>
                        <option selected disabled>-- Docker Image --</option>`;
            optoo = `<option id="`;
            optoc = `">`;
            optc = `</option>`;
            footer = `</select>`;
            
            var body = "";

            for (let i = 0; i < imgs.length; i++) {
               id = imgs[i].ID;
               name = imgs[i].Name;
               repo = imgs[i].Repo;
               
               body = body + optoo + id + optoc + name + " | " + repo + optc;
            }

            ddm = header + body + footer;

            document.getElementById("image-ddm").innerHTML = ddm;

         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }

};

/*
 * ENVIRONMENTS
 */

function launchAnimate() {

   var launch1 = `<span id="launch-env-1" class="loading"> ... </span>`;
   var launch2 = `<span id="launch-env-2" class="loading"> ... </span>`;

   document.getElementById("launch-env-1").innerHTML = launch1;
   document.getElementById("launch-env-2").innerHTML = launch2;

}

function launchRestore() {

   var launch1 = `<span id="launch-env-1"><img src="./assets/images/launch-scaled.svg" alt=""></span>`;
   var launch2 = `<span id="launch-env-2"> LAUNCH </span>`;
   document.getElementById("launch-env-1").innerHTML = saving;
   document.getElementById("launch-env-2").innerHTML = saving;

}

// Add environment
window.addEnv = function(image, name, hostname, port, local, path, options) {
   var data = {};
  
   if (image == "" || name == "" || hostname == "" || local == ""  || path == "") {
      launchRestore();
      document.getElementById("add-env-form").reset();
      return
   }

   try {
      window.go.main.App.AddEnvironment(image, name, hostname, port, local, path, options)
         .then(result => {
            data = result;
            document.getElementById("tab-4").checked = true;
            getEnvs();
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }

   document.getElementById("add-env-form").reset();
};

function saveAnimate() {

   var saving1 = `<span id="edit-env-save-1" class="loading"> ... </span>`;
   var saving2 = `<span id="edit-env-save-2" class="loading"> ... </span>`;

   document.getElementById("edit-env-save-1").innerHTML = saving1;
   document.getElementById("edit-env-save-2").innerHTML = saving2;

}

function saveRestore() {

   var saving1 = `<span id="edit-env-save-1" class="loading"><img src="./assets/images/save-scaled.svg" alt=""></span>`;
   var saving2 = `<span id="edit-env-save-2" class="loading"> SAVE </span>`;

   document.getElementById("edit-env-save-1").innerHTML = saving1;
   document.getElementById("edit-env-save-2").innerHTML = saving2;

}

// Change environment function
window.changeEnv = function(id, image, name, hostname, port, local, path, options) {
   var data = {};
   
   if (image == "" || name == "" || hostname == "" || local == ""  || path == "") {
      saveRestore();
      document.getElementById("edit-env-form").reset();
      return
   }

   try {
      window.go.main.App.ChangeEnvironment(id, image, name, hostname, port, local, path, options)
         .then(result => {
            data = result;
            saveRestore();

            getEnvs();
            document.getElementById("tab-4").checked = true;
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }

     document.getElementById("edit-env-form").reset();

};

// Remove environment
window.removeEnv = function(id) {
   var data = {};
   
   btn = "ss-btn-" + id;
   document.getElementById(btn).innerHTML = `<span class="loading"> ... </span>`;

   try {
      window.go.main.App.RemoveEnvironment(id)
         .then(result => {
            data = result;
            getEnvs();
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
};

// Get all images for dropdown menu and preselect existing value
window.getImgDdmSelected = function(selected) {
   
   try {
      data = window.go.main.App.GetImages()
         .then(result => {
            imgs = result;

            header = `<select id="env-edit-image" required>
                        <option disabled>-- Docker Image --</option>`;
            optoon = `<option id="`;
            optoos = `<option selected id="`;
            optoc = `">`;
            optc = `</option>`;
            footer = `</select>`;
            
            var body = "";

            for (let i = 0; i < imgs.length; i++) {
               id = imgs[i].ID;
               name = imgs[i].Name;
               repo = imgs[i].Repo;

               var optoo = optoon;
               compare = name + " | " + repo;
               if (compare == selected) {
                  optoo = optoos;
               }
               
               body = body + optoo + id + optoc + name + " | " + repo + optc;
            }

            ddm = header + body + footer;

            document.getElementById("env-image-ddm").innerHTML = ddm;

         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }

};

// Get single environment 
window.getEnv = function(id) {

   try {
      data = window.go.main.App.GetEnvironment(id)
         .then(result => {
            env = result;
            var id = env.ID;
            var image = env.Image;
            var name = env.Name;
            var hostname = env.Hostname;
            var port = env.Port;
            var local = env.Local;
            var path = env.Path;
            var options = env.Options;
                        
            getImgDdmSelected(image);

            eef11 = `<input type="text" id="env-edit-name" placeholder="Container Name" value="` + name + `" required>`;  
            eef12 = `<input type="text" id="env-edit-hostname" placeholder="Container Hostname" value="` + hostname + `" required>`;
            eef13 = `<input type="text" id="env-edit-port" placeholder="Port:Port" value="` + port + `" required>`;  

            var eef1 = eef11 + eef12 + eef13;

            document.getElementById("env-edit-form-1").innerHTML = eef1;

            eef21 = `<input type="text" id="env-edit-local" placeholder="Local Code Repo" value="` + local + `" required>`;
            eef22 = `<input type="text" id="env-edit-path" placeholder="Container Path" value="` + path + `" required>`;
            eef23 = `<input type="text" id="env-edit-id" class="hidden" value="` + id + `" required>`;

            var eef2 = eef21 + eef22 + eef23;

            document.getElementById("env-edit-form-2").innerHTML = eef2;

            eef3 = `<input type="text" id="env-edit-options" value="` + options + `">`;

            document.getElementById("env-edit-form-3").innerHTML = eef3;

            document.getElementById("tab-4.5").checked = true;
         })
         .catch((err) => {
            console.error(err);
         });
      } catch (err) {
         console.error(err);
      }
};

// Access single environment 
window.accessEnv = function(id, name) {

   var data = {};

   btn = "ss-btn-" + id;
   console.log("Button Log");
   console.log(btn);
   state = document.getElementById(btn).innerHTML;

   if (state == "start") {
      call_popup("You have to first start the environment before you can access it.", 0, "white", "black");
      return
   }

   try {
      window.go.main.App.AccessEnvironment(name)
         .then(result => {
            data = result;
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
};

// Access logs of single environment 
window.accessLogs = function(id, name) {

   var data = {};

   btn = "ss-btn-" + id;
   console.log("Button Log");
   console.log(btn);
   state = document.getElementById(btn).innerHTML;

   if (state == "start") {
      call_popup("You have to first start the environment before you can access it.", 0, "white", "black");
      return
   }

   try {
      window.go.main.App.AccessLogs(name)
         .then(result => {
            data = result;
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
};

// Get all environments
window.getEnvs = function() {

   try {
      data = window.go.main.App.GetEnvironments()
         .then(result => {
            envs = result;
            header = `<table id="env-table">
                   <thead>
                     <tr>  
                       <th>name <button onclick="sortTable(0, 'env-table')" class="t-nav"><img src="./assets/images/d-arrow.png"></button></th>
                       <th>ip address</th>
                       <th>ports</th>
                       <th class="t-center">action</th>
                       <th class="t-center"></th>
                       <th class="t-center"></th>
                       <th class="t-center">edit</th>
                       <th class="t-center"></th>
                     </tr>
                   </thead>
                   <tbody>`;
            part1 = `<tr>
                       <td width="150">`;
            part2 = `  </td>
                        <td width="150">`;
            part3 = `</td>
                     <td width="150">`;

            part4 = `</td>
                     <td class="t-center"><button id="ss-btn-`;

            part40 = `" onclick="`;
            
            part41 = `')" class="table-btn">`;
            
            part42 = `</button></td>`;

            part5 = `<td class="t-center"><button id="access-`;

            part51 = `" onclick="accessEnv('`;

            part52 = `')" class="table-btn access-btn"><img class="access-icon" src="./assets/images/terminal.png"></button></td><td class="t-center"><button id="logs-`;

            part53 = `" onclick="accessLogs('`;

            part54 = `')" class="table-btn logs-btn"><img class="logs-icon" src="./assets/images/logs-yellow.png"></button></td>
                     <td class="t-center"><button onclick="editEnv('`; 
            
            part6 = `')" class="table-btn">edit</button></td>
                       <td class="t-center"><button onclick="removeEnv('`;
            part7 = `')" id="rm-`;

            part8 = `" class="table-btn remove-btn"><img class="remove-icon" src="./assets/images/remove.png"></button></td>
                     </tr>`;
                     //<tr class="spacer"><td colspan="100"></td></tr>`;
            footer =`
                     </tbody>
                        </table>
                        <button onclick="checkTab('tab-5')" class="add-btn">+</button>`;

            var body = "";

            for (let i = 0; i < envs.length; i++) {
               id = envs[i].ID;
               name = envs[i].Name;
               ip = envs[i].IP;
               port = envs[i].Port;
               action = envs[i].Action;

               method = "";
               if (action == "start") {
                  method = `startEnv('`;
               }

               if (action == "stop") {
                  method = `stopEnv('`;
               }

               body = body + part1 + name + part2 + ip + part3 + port + part4 + id + part40 + method + id + part41 + action + part42 + part5 + id + part51 + id + `', '` + name + part52 + id + part53 + id + `', '` + name + part54 + id + part6 + id + part7 + id + part8;
            }

            var table = header + body + footer;

            document.getElementById("environment-table").innerHTML = table;

         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }

};

// Start environment
window.startEnv = function(id) {
   var data = {};

   btn = "ss-btn-" + id;
   document.getElementById(btn).innerHTML = `<span class="loading"> ... </span>`;
   
   try {
      window.go.main.App.StartEnv(id)
         .then(result => {
            data = result;
            document.getElementById("tab-4").checked = true;
            getEnvs();
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
};

// Stop environment
window.stopEnv = function(id) {
   var data = {};
   
   document.getElementById("ss-btn-" + id).innerHTML = `<span class="loading"> ... </span>`;
   
   try {
      window.go.main.App.StopEnv(id)
         .then(result => {
            data = result;
            document.getElementById("tab-4").checked = true;
            getEnvs();
         }) 
         .catch((err) => {
           console.error(err);
         });
     } catch (err) {
       console.error(err);
     }
};

