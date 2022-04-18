/* Do not change, this code is generated from Golang structs */

export {};

export class Environment {


    static createFrom(source: any = {}) {
        return new Environment(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class EnvInJSON {
    ID: number;
    Image: string;
    Name: string;
    Hostname: string;
    IP: string;
    Port: string;
    Local: string;
    Path: string;
    Options: string;
    Action: string;
    Status: string;

    static createFrom(source: any = {}) {
        return new EnvInJSON(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.ID = source["ID"];
        this.Image = source["Image"];
        this.Name = source["Name"];
        this.Hostname = source["Hostname"];
        this.IP = source["IP"];
        this.Port = source["Port"];
        this.Local = source["Local"];
        this.Path = source["Path"];
        this.Options = source["Options"];
        this.Action = source["Action"];
        this.Status = source["Status"];
    }
}
export class Image {


    static createFrom(source: any = {}) {
        return new Image(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class ImageInJSON {
    ID: number;
    Name: string;
    Repo: string;

    static createFrom(source: any = {}) {
        return new ImageInJSON(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.ID = source["ID"];
        this.Name = source["Name"];
        this.Repo = source["Repo"];
    }
}
export class Table {


    static createFrom(source: any = {}) {
        return new Table(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}