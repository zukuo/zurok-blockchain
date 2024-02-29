export namespace gui {
	
	export class balances {
	    key: number;
	    address: string;
	    balance: number;
	
	    static createFrom(source: any = {}) {
	        return new balances(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.address = source["address"];
	        this.balance = source["balance"];
	    }
	}
	export class blocks {
	    key: number;
	    hash: string;
	    prevhash: string;
	    height: number;
	    timestamp: string;
	    nonce: number;
	    pow: boolean;
	
	    static createFrom(source: any = {}) {
	        return new blocks(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.hash = source["hash"];
	        this.prevhash = source["prevhash"];
	        this.height = source["height"];
	        this.timestamp = source["timestamp"];
	        this.nonce = source["nonce"];
	        this.pow = source["pow"];
	    }
	}
	export class txout {
	    value: number;
	    pubkeyhash: string;
	
	    static createFrom(source: any = {}) {
	        return new txout(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.pubkeyhash = source["pubkeyhash"];
	    }
	}
	export class txin {
	    txid: string;
	    vout: number;
	    signature: string;
	    pubkey: string;
	
	    static createFrom(source: any = {}) {
	        return new txin(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.txid = source["txid"];
	        this.vout = source["vout"];
	        this.signature = source["signature"];
	        this.pubkey = source["pubkey"];
	    }
	}
	export class transactions {
	    key: number;
	    transaction: string;
	    amount: number;
	    block: string;
	    height: number;
	    inputs: txin[];
	    outputs: txout[];
	
	    static createFrom(source: any = {}) {
	        return new transactions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.transaction = source["transaction"];
	        this.amount = source["amount"];
	        this.block = source["block"];
	        this.height = source["height"];
	        this.inputs = this.convertValues(source["inputs"], txin);
	        this.outputs = this.convertValues(source["outputs"], txout);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

