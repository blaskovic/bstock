# bstock
Just another stock portfolio

**Input yaml file**

    stocks:
      BAASTOCK:
        url: http://www.bcpp.cz/Cenne-Papiry/Detail.aspx?isin=GB00BF5SDZ96
        notes: Will go up
        currency: CZK
        buyprice: 74.05
        amount: 100
        fees: 100
      NYSE:UPL:
        url: https://www.google.com/finance?q=NYSE:UPL&ei=TsXYVJDHLKmewwOhyoC4DQ
        notes: Some note
        currency: USD
        buyprice: 13.40
        amount: 50
        fees: 15.90

**Output**

    +--------------+-----------+-----------+----------+--------+------------+-----------+- ----------+
    |    TICKER    |   PRICE   | BUY PRICE |   DIFF   | DIFF % |    FEES    |  OVERALL  |    NOTES   |
    +--------------+-----------+-----------+----------+--------+------------+-----------+------------+
    | 100 BAASTOCK | 75.50 CZK | 74.05 CZK | 1.45 CZK | 0.02 % | 100.00 CZK | 45.00 CZK | Will go up |
    |  50 NYSE:UPL | 14.83 USD | 13.40 USD | 1.43 USD | 0.10 % |  15.90 USD | 55.60 USD |  Some note |
    +--------------+-----------+-----------+----------+--------+------------+-----------+------------+
    
