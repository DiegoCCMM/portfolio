# Fullstack project

The idea behind this project is to condense all the knowledge i've gathered as a full stack developer this years.

# Objectives
This project shows my ability to create from zero a whole FE & BE flow and consuming data sources for later usage

# The how
The project consists in consuming https://www.thecocktaildb.com/api.php an API how cocktails recipes and data

This data will be dumped, using Python, in a non-relation DB such as MongoDB
Then an API will consume this DB and bring the FE everything it will need.
Lastly, a simple FE will print in angular the data retrieved from the BE.


# Development
Data dump is done by the  E and the TL python scripts that uses the JSON format the source gives us to dump it directly on the mongoDB
Using mongo compass, connected to the localhost process running the DB, we can see all the diferent documents under the database cocktail_db and the collection cocktails.

