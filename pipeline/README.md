# Weather pipeline

This can be run on a server to collect weather data (i.e. schedule a cron job to run the executable). 

Requires a SQLite database named db.sqlite in the same directory as the executable. To create this database and the 
required tables, run <code>initdb.sh</code>.

Also requires a <code>secrets.json</code> file in the same directory as the executable with three keys: 
- apiKey - an open weather map API key
- lat - latitude
- lon - longitude

### Compile
<code>make</code>

### Run
<code>make run</code>
