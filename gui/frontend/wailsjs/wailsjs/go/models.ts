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

}

