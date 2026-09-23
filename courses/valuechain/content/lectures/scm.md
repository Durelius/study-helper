---
id: scm
title: "Supply Chain Management"
deck: "new 03"
slides: 40
summary: "How value moves once it leaves one company: the actors and the network, demand planning, inventory, push vs pull, sourcing and logistics, and the bullwhip effect."
---

## Where supply chain sits in the course

The deck is the third week of a three-week arc: **Week 1 Value Chain** asks *where is value created?*, **Week 2 Operations** asks *how is value created efficiently?*, and **Week 3 Supply Chain** asks *how do organisations coordinate flows externally?*

That last word is the whole distinction. **Operations management** manages flow **within one organisation** — the deck's example is a coffee order: Order → Pay → Prepare → Pickup, one company, one process. **Supply chain management** coordinates flow **across organisations** — Supplier → Manufacturer → Distributor → Retailer → Customer. Zoom out from the process and you get the chain.

A **supply chain (SC)** is defined as a **flow of materials, information and money** across suppliers, manufacturers and customers — multiple entities. **SCM** is managing that flow **efficiently and strategically**. Note that the arrow on the slide points both ways: money and information run back up the chain while materials run down it.

`[slides 2-5]`

## Actors, enablers, and why it is a network

| Actor | Role | What it is judged on |
|---|---|---|
| Suppliers | Provide raw materials, parts, components, services | Quality, on-time delivery, supporting innovation |
| Manufacturers | Transform inputs into finished products | Efficient production, quality & safety, capacity & cost |
| Distributors | Move and store products, bridge supply and demand | Consolidate inventory, availability, right place |
| Wholesalers | Buy in bulk, sell to business customers | Hold inventory, product variety, market reach |
| Retailers | Sell to end customers | Convenience, service, brand value |
| Customers | Buy and use products | Drive demand, give feedback, build loyalty |

Three **enablers** make the chain work: **logistics** (right product, right place, right time), **technology** (information systems, data visibility, automation, collaboration platforms) and **finance** (working capital, trade finance, risk management, payment solutions).

Two structural points follow. First, the textbook straight line is a simplification: **real supply chains are networks**, with multiple suppliers, multiple factories, multiple 3PL providers, multiple channels and multiple customer segments all cross-linked. Second, **one company can occupy several roles at once** — CP Group is simultaneously supplier, manufacturer, distributor/logistics operator and retailer (7-Eleven, Lotus's, Makro), and it runs *several distinct product supply chains* (animal feed, meat, seafood, ready-to-eat, dairy, pharma) inside one group.

`[slides 6-9]`

## Supply chain performance and its trade-offs

Good internal processes are necessary but not sufficient — **the entire network must align**. The deck maps Week 2's process objectives onto supply-chain-level (SCOR) attributes:

| Process level (Week 2) | Supply chain level (Week 3) | Meaning |
|---|---|---|
| Quality · Speed · Dependability · Flexibility · Cost | **Reliability** | Deliver what was promised, in full, on time, in perfect condition |
| | **Responsiveness** | Respond quickly to customer needs and demand changes |
| | **Agility** | Adapt and reconfigure the supply chain rapidly |
| | **Cost** | Minimise total supply chain cost while creating value |
| | **Asset efficiency** | Maximise utilisation of assets across the chain |

No chain maximises everything. The three **strategic trade-offs**, each with its example KPIs:

- **Low cost ↔ high responsiveness** — cost per unit (lower is better) vs order cycle time (shorter is better)
- **Low inventory ↔ high availability** — inventory turns (higher is better) vs fill rate (higher is better)
- **Efficiency ↔ flexibility** — operating cost ratio (lower is better) vs time to adapt (shorter is better)

`[slides 10-14]`

## Efficient vs responsive, and strategic fit

| | **Efficient** | **Responsive** |
|---|---|---|
| Demand | Predictable | Uncertain |
| Priority | Cost | Speed / flexibility |
| Inventory | Lean | Buffer |
| Capacity | High utilisation | Capacity cushion |
| Suppliers | Cost-focused | Flexible |
| Example | Commodity | Fashion |

The deck later widens this into four strategy types: **efficient** (cost & utilisation; stable demand; Walmart, Big C, Tesco), **responsive** (speed & availability; uncertain demand, short life cycles; Zara, Uniqlo), **agile** (flexibility & adaptability; high uncertainty and variety; small batches, flexible suppliers, postponement; Louis Vuitton, Dell) and **hybrid / leagile** (balance both; *push upstream, pull downstream*, place the push–pull boundary, segment customers and products).

> **Exam focus.** The **strategic fit** case — Big C vs Zara — is the most likely case-study prompt in this deck. **Big C**: cost leadership, predictable staple demand, **push** (produce and distribute in advance), efficiency and scale; KPIs cost per unit, inventory turns, on-time in-full, supply chain cost %. **Zara**: differentiation, trend-driven uncertain demand, **pull-hybrid** (design → make → deliver on real demand signals), speed and flexibility; KPIs lead time to store, sell-through rate, stockout rate, new product contribution. The examiner wants strategy → demand type → SC approach → KPI, lined up consistently.

`[slides 15-16, 27]`

?check id=scm-010

## Demand drives everything

Supply-chain decisions cascade from one input: **customer demand → forecast / orders → inventory → production → sourcing → transportation**. Higher demand needs more forecast and orders, which need more inventory (or risk a stockout), which needs more production capacity, which needs more materials, which needs more transport.

**Demand planning** estimates future customer demand to support those decisions. Inputs: historical sales, customer orders, promotions, seasonality, market trends. Process: cross-functional teams combine, analyse and discuss — collaborate, apply business knowledge, challenge assumptions, agree **one demand view**. Output: a **demand plan**, a time-phased view of expected demand by volume (what), timing (when), location (where) and segment/channel (who). The deck's caution: it is *not about perfect prediction, it is about the best informed plan*.

Forecasts are necessary but never perfectly accurate, and the two errors are not symmetric:

| | **Underforecast** (forecast 100, actual 120) | **Overforecast** (forecast 100, actual 80) |
|---|---|---|
| Gap | 20 units unmet demand | 20 units excess inventory |
| Result | **Stockout** — lost sales, dissatisfied customers, expedited shipments, higher costs | **Excess inventory** — higher holding costs, obsolescence risk, markdowns, cash tied up |

`[slides 17-19]`

## Inventory: why, where, how much

Companies hold inventory to **meet customer demand**, **buffer uncertainty**, **decouple activities** and **gain purchasing/production efficiency**.

The target is not zero and not maximum but **the right inventory** — the balance of service level against total cost. **Too little**: stockouts, lost sales, poor service (long waits, backorders). **Too much**: holding cost, obsolescence, waste (damage, spoilage, shrinkage) and **capital tied up**.

Inventory sits at four places, each with a different job:

| Stage | Type | Purpose |
|---|---|---|
| Supplier | **Raw materials** | Reliable supply, buffer supplier variability, economies of scale |
| Factory | **WIP (work-in-process)** | Keep production flowing, balance process times, absorb process variability |
| Warehouse | **Finished goods** | Meet demand, smooth demand variability, responsive delivery |
| Retailer | **Retail inventory** | High availability, variety and choice, drive sales |

`[slides 20-22]`

## Push, pull and the push–pull boundary

**Push** is triggered by **forecast**: make-to-stock, produce in advance, inventory built **before** demand occurs. Best when demand is predictable, lead times are long and economies of scale matter (canned food, beverages, paper, household products).

**Pull** is triggered by **actual customer demand or order**: make-to-order, receive order → produce → ship (or assemble to order), inventory built **after** demand is known or kept minimal. Best when demand is uncertain or variable, customisation is needed, and overproduction risk must be low (custom furniture, made-to-order appliances, engineered equipment).

Most real chains are **hybrid**. The **push–pull boundary** is the point where the chain switches from forecast-driven to demand-driven: raw materials, WIP and finished goods are pushed to the warehouse, then **postponement and final customisation** happen on actual demand. Where you place the boundary depends on **demand uncertainty, product variety, lead time, and cost/service strategy**.

`[slides 23-24]`

?check id=scm-016

## Network design, sourcing and logistics

**Network design** answers four questions: how many **suppliers** (single vs multiple, global vs local, capability and reliability); how many and where the **factories** (capacity, technology, flexibility); how many and where the **warehouses** (inventory positioning strategy); and which **distribution channels** (direct vs indirect, omnichannel integration).

**Sourcing strategy** has two dimensions and a trade-off on each:

| | **Single sourcing** | **Multiple sourcing** |
|---|---|---|
| Cost | Lower unit cost | Higher cost |
| Complexity | Simpler management | More complex |
| Supply risk | **High** | **Lower** |
| Flexibility | Less | More |
| Buying power | Weaker leverage | Stronger leverage |

| | **Global sourcing** | **Local / nearshore sourcing** |
|---|---|---|
| Cost | Lower unit cost | Higher unit cost |
| Lead time | Longer | Shorter |
| Supply risk | Greater (disruption, tariffs) | Lower — faster and more resilient |
| Capability | Broader capabilities | Limited capabilities / capacity |
| Sustainability | Higher emissions | Lower emissions |

**Logistics** manages the movement and storage of goods across the chain, through four activities: **transportation** (road, rail, air, sea, pipeline — cost, speed, reliability), **warehousing** (capacity, location, cost), **order fulfilment** (pick, pack, ship — accuracy, speed, service) and **information / tracking** (visibility, status, coordination).

| Mode | Cost | Speed | Best for |
|---|---|---|---|
| Air | High | Fast (2–5 days) | High-value / urgent, perishables, e-commerce |
| Sea | Low | Slow (25–40+ days) | High-volume / global, bulk, low-value, heavy |
| Road | Medium | Flexible | Regional |
| Rail | Low–medium | Medium | Large inland flows |

Air is more predictable on schedule but high emissions per kg; sea is more variable (weather, port congestion) but low emissions per kg. The deck's rule: **there is no "best" mode — only the best choice for your priorities** (cost → sea, speed → air, reliability → air, sustainability → sea).

Facility count is the same kind of trade-off. **Centralised**: fewer facilities, lower inventory, economies of scale, longer delivery. **Decentralised**: more facilities, higher inventory, closer to the customer, faster delivery.

`[slides 25-31, 34]`

## The bullwhip effect

**Bullwhip effect**: small changes in customer demand create **progressively larger fluctuations upstream**. Actual customer demand is relatively stable; retailer orders to the wholesaler are more variable; wholesaler orders to the distributor more variable still; distributor orders to the factory highly variable; factory production extremely variable. Small variation downstream, large variation upstream.

Five causes:

| Cause | Mechanism |
|---|---|
| **Forecast updating** | Each stage re-forecasts independently on limited information and overreacts to recent changes |
| **Order batching** | Large, infrequent orders instead of small frequent ones create artificial demand spikes upstream |
| **Promotions** | Sales promotions cause temporary surges that upstream stages misread as real growth |
| **Long lead times** | Longer lead times raise uncertainty, encouraging bigger safety stocks and bigger orders |
| **Poor information sharing** | Limited visibility of real customer demand and inventory means stages guess instead of using facts |

The countermeasure the deck gives is **visibility and coordination**. **Real-time dashboards** show live inventory, transit times and sales, enabling data-driven decisions in real time and directly attacking the poor-information-sharing cause. The wrap-up states it plainly: *coordination and visibility reduce bullwhip and improve performance*. The cases: **Zara** pushes real-time sales data from stores to HQ, runs fast design and reorder cycles and decentralises decision-making; **Amazon** uses AI-based inventory prediction, live fulfilment dashboards and optimised last-mile delivery.

> **Exam focus.** Bullwhip is the highest-yield item in this deck. Be able to state the definition (amplification *upstream*, not downstream), list the five causes, and match a one-line scenario to the right cause — a full-truckload order policy is *order batching*; a discount week misread as growth is *promotions*; a six-week ocean lead time driving bigger safety stock is *long lead times*.

`[slides 32-33, 38-40]`

?check id=scm-023

## The Beer Game simulation

Each team is one beer supply chain with four roles: **Retailer** (sells to customers) → **Wholesaler** (supplies the retailer) → **Distributor** (supplies the wholesaler) → **Manufacturer** (produces for the distributor).

Rules: the game runs **20 weeks (rounds)**; **initial inventory is 10/10/10 cases**. Each week a player receives the shipment from upstream and production, checks and fills incoming orders, records inventory or backorders, and places a new order upstream. **No talking about demand forecasts** outside your orders — the rule that manufactures the poor-information-sharing condition.

Costs, and the only arithmetic in the deck:

- **Inventory: $0.50 per case, per week**
- **Backlog (unmet demand): $1.00 per case, per week**
- Closing position each week = (inventory + production/receipt) − units shipped. The worked row on the slide: start 10, receive 10 → 20 before shipping; ship A → close at **20 − A**.

The team with the **lowest total cost** wins — note that backlog is priced at twice inventory, so stocking out is the expensive error. The debrief asks what happened to inventory levels, where the biggest delays occurred, how communication gaps affected decisions, and how this resembles real business — i.e. it is a bullwhip demonstration.

`[slides 35-37]`

## Wrap-up

Demand drives supply-chain decisions. Inventory balances availability against cost. Strategy determines whether you prioritise efficiency or responsiveness. Network, sourcing and logistics determine how supply is positioned and moved. Coordination and visibility reduce bullwhip and improve performance.

`[slide 40]`
