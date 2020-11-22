# Stashaway Code Challenge
author: Daniel Silva


## Assumptions
for this challenge, a few assumtions were made

- For SingleTime deposit plans, only the first deposit will be included. everything else won't be.
- For Montly plans, only the first deposit of the month will be considered. 
- For sake of simplicity, incoming deposit will be distributed using this logic:
    + funds will be added using a greedy approach. in other words, portolios with higher amounts will be funded first than lower amounts.

- There will be a default account. all additional money what does not fit on any portfolio, will be sent there. the idea will be that, customer could manually assign it to any porfolio later 
- a user sent deposits once a month