package bank;

import java.util.HashSet;
import java.util.Set;

import account.Account;

public class Bank {

    private Set<Account> accounts=new HashSet<Account>();

    public void add(Account account) {
		accounts.add(account);
    }

    public void accrue(double rate) {
		for (Account account: accounts) {
			account.accrue(rate);
		}
    }

    public String toString() {
		String string = "";
		for (Account account: accounts) {
			string += account + "\n";
		}
		return string ;
    }
}
