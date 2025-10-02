export namespace main {
	
	export class PrintJobDisplay {
	    id: string;
	    file_extension: string;
	    printer_name: string;
	    // Go type: time
	    created: any;
	    status: string;
	    messages: string[];
	
	    static createFrom(source: any = {}) {
	        return new PrintJobDisplay(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.file_extension = source["file_extension"];
	        this.printer_name = source["printer_name"];
	        this.created = this.convertValues(source["created"], null);
	        this.status = source["status"];
	        this.messages = source["messages"];
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
	export class AllJobs {
	    current_jobs: PrintJobDisplay[];
	    recent_jobs: PrintJobDisplay[];
	
	    static createFrom(source: any = {}) {
	        return new AllJobs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current_jobs = this.convertValues(source["current_jobs"], PrintJobDisplay);
	        this.recent_jobs = this.convertValues(source["recent_jobs"], PrintJobDisplay);
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
	export class NetworkDevice {
	    already_active: boolean;
	    hostname: string;
	    host: string;
	    port: string;
	
	    static createFrom(source: any = {}) {
	        return new NetworkDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.already_active = source["already_active"];
	        this.hostname = source["hostname"];
	        this.host = source["host"];
	        this.port = source["port"];
	    }
	}
	
	export class RecentScan {
	    created: string;
	    code: string;
	    result: string;
	
	    static createFrom(source: any = {}) {
	        return new RecentScan(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.created = source["created"];
	        this.code = source["code"];
	        this.result = source["result"];
	    }
	}
	export class RemoteConnectionData {
	    workstation_name: string;
	    url: string;
	    last_ping: string;
	
	    static createFrom(source: any = {}) {
	        return new RemoteConnectionData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.workstation_name = source["workstation_name"];
	        this.url = source["url"];
	        this.last_ping = source["last_ping"];
	    }
	}
	export class SubnetSearch {
	    results: NetworkDevice[];
	    search_term: string;
	
	    static createFrom(source: any = {}) {
	        return new SubnetSearch(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.results = this.convertValues(source["results"], NetworkDevice);
	        this.search_term = source["search_term"];
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
	export class USBDevice {
	    id: string;
	    vendor_id: string;
	    product_id: string;
	    name: string;
	    active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new USBDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.vendor_id = source["vendor_id"];
	        this.product_id = source["product_id"];
	        this.name = source["name"];
	        this.active = source["active"];
	    }
	}

}

export namespace printer {
	
	export class StatusChangeRequest {
	    id: string;
	    messages: string[];
	
	    static createFrom(source: any = {}) {
	        return new StatusChangeRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.messages = source["messages"];
	    }
	}
	export class PrintClientPing {
	    printers: printers.Printer[];
	    label_id: string;
	    cancel_print_jobs: StatusChangeRequest[];
	    success_print_jobs: StatusChangeRequest[];
	
	    static createFrom(source: any = {}) {
	        return new PrintClientPing(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.printers = this.convertValues(source["printers"], printers.Printer);
	        this.label_id = source["label_id"];
	        this.cancel_print_jobs = this.convertValues(source["cancel_print_jobs"], StatusChangeRequest);
	        this.success_print_jobs = this.convertValues(source["success_print_jobs"], StatusChangeRequest);
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

export namespace printers {
	
	export class Printer {
	    id: string;
	    name: string;
	    active: boolean;
	    network: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Printer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.active = source["active"];
	        this.network = source["network"];
	    }
	}

}

