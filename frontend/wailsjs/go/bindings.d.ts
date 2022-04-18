import * as models from './models';

export interface go {
  "main": {
    "App": {
		AccessEnvironment(arg1:string):Promise<boolean>
		AccessLogs(arg1:string):Promise<boolean>
		AddEnvironment(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string,arg6:string,arg7:string):Promise<boolean|string>
		AddImage(arg1:string,arg2:string):Promise<void>
		ChangeEnvironment(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string,arg6:string,arg7:string,arg8:string):Promise<boolean>
		ChangeImage(arg1:string,arg2:string,arg3:string):Promise<boolean>
		GetEnvironment(arg1:string):Promise<models.Environment>
		GetEnvironments():Promise<Array<models.EnvInJSON>|boolean>
		GetImage(arg1:string):Promise<models.Image>
		GetImages():Promise<Array<models.ImageInJSON>|boolean>
		Greet(arg1:string):Promise<string>
		RemoveEnvironment(arg1:string):Promise<boolean>
		RemoveImage(arg1:string):Promise<void>
		StartEnv(arg1:string):Promise<boolean>
		StopEnv(arg1:string):Promise<boolean>
		Test():Promise<models.Table>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
