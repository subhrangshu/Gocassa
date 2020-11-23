# Gocassa
OpenSource ORM Library for Apache Cassandra Distributed NoSQL Databases with GO | GOLang

### This library is currently in building phase, if you want to contribute to the project, fork the repo and make pull requests.

To get going, install GO to your system.
For linux, in terminal type ```snap install go --classic``` use ```sudo``` if required.

Install Apache Cassandra
```sudo apt-get install cassandra```
If this gives error, follow official apache cassandra site to add repository to PPA and rerun the above command.
You might also reconfigure `cassandra-env.sh` file to retrict RAM and CPU usage else JVM heap memory might start clogging your system with unnecessary load.

Install `cqlsh` by `pip install cqlsh` (keep in mind this works only with python2 and in case this still gives hurdles, use `conda`, `pip` environments or `docker`, `podman` container.

After `cqlsh` installation, in python2 env linux terminal type `cqlsh'` to start `cqlsh`. If this gives version mismatch error, use command `cqlsh --cqlversion=3.4.5` to bypass the error. (Change version number if the interpreter asks so)

In `> cqlsh`, create a keyspace by, 
```
CREATE KEYSPACE test_keyspace
  WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
  };
```
  
Now in your linux pc, go to a suitable directory and clone the repository with `git clone https://github.com/subhrangshu/Gocassa ` this shall clone the repo, and if you can't find `git` in your system, install it by `sudo apt-get install git`

Next `$ cd Gocassa' to move into the newly created directory.

To install `gocql` driver, use the command `$ go get -u github.com/gocql/gocql`

Finally run `$ go run main.go` to compile and run the library.


In case you find difficulty, contact me directly on WhatsApp on `+91-977-547-9892`


Subhrangshu Adhikary
