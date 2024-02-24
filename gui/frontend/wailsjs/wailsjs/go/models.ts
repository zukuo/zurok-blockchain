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

}

