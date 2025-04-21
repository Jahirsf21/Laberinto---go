export namespace main {
	
	export class Coord {
	    fila: number;
	    col: number;
	
	    static createFrom(source: any = {}) {
	        return new Coord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fila = source["fila"];
	        this.col = source["col"];
	    }
	}
	export class Maze {
	    name: string;
	    mode: boolean;
	    matrix: number[][];
	    start: Coord;
	    end: Coord;
	
	    static createFrom(source: any = {}) {
	        return new Maze(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.mode = source["mode"];
	        this.matrix = source["matrix"];
	        this.start = this.convertValues(source["start"], Coord);
	        this.end = this.convertValues(source["end"], Coord);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MazeData {
	    Mazes: Maze[];
	
	    static createFrom(source: any = {}) {
	        return new MazeData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mazes = this.convertValues(source["Mazes"], Maze);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

