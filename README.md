dalo - the daily achievement logger
===================================

Use this small quick and dirty command line tool to log your daily achievements.
Data is stored in a JSON file. Tested on Arch Linux with Go 1.4.

Setup
-----
1. `go get -u github.com/hoffie/dalo`
2. `export DALO_DB=~/.dalo.db` (I've put something like this in my .bashrc)

Usage
-----
  * `dalo Created a tool for tracking daily achievements` inserts a new entry for today
  * `dalo 2015-05-04 Successfully logged an entry for a different date` logs an entry for a date other than today
  * `dalo` displays all entries per date in a human-readable format, sorted from older to newer
  
    ```
    2015-05-04
    ==========
    * Successfully logged an entry for a different date

    2015-05-05
    ==========
    * Created a tool for tracking daily achievements
    ```
  * `dalo 2015-05-05` shows all entries for the given date
  
    ```
    2015-05-05
    ==========
    * Created a tool for tracking daily achievements
    ```
    
No other functionality is currently implemented. Really.
