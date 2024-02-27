# blackbox-rebalance-service

Micro-service for rebalancing portfolios in my blackbox application.

When my trade signal service (private) updates a portfolio table, a postgres trigger fires a rebalance request to this service with a strategyId. This service looks for all accounts that are subscribed to that strategy, checks their holdings compared to the updated portfolio table, and makes the necessary trades to balance each users portfolio with strategy they are following.

hosted on fly.io
use flyctl deploy to deploy the app to fly

use air to run this locally
https://github.com/cosmtrek/air
