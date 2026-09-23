---
id: logistics
title: "Supply Chain Strategy & Global Logistics"
deck: "old 03"
slides: 33
summary: "How a firm's value-chain strategy determines order fulfillment, how supply chains are structured as stakeholder networks, and why weak coordination across that network amplifies into the bullwhip effect."
---

## Recap: from value-chain strategy to order fulfillment

The deck opens by tying this week back into the value-chain vs. business-process distinction: the **value chain** is the broad, strategic view of what activities create value; the **business process map** zooms into how one high-impact activity actually works, step by step.

Two generic strategies are re-examined through the lens of **order fulfillment** (outbound logistics):

| Strategy | Focus | Order fulfillment priority | Business-process features |
|---|---|---|---|
| **Differentiation** | Customer experience and responsiveness | Speed, accuracy, flexibility — fast reliable delivery becomes part of the brand promise and can justify premium pricing | Real-time tracking & proactive updates, priority picking & quality checks, multiple delivery options |
| **Cost Leadership** | Efficiency | Efficient, low-cost fulfillment that minimizes waste, labor cost, and shipping errors | Automated picking/packing, batch processing for speed, standardized packaging & optimized routes |

The deck's **Devil's Quadrangle** makes the underlying trade-off explicit: a firm cannot maximize **Cost, Speed (Time), Flexibility, and Quality** performance all at once. On the radar chart, a Cost-Leadership shape stretches furthest toward Cost performance, while a Differentiation shape stretches toward Quality, Time, and Flexibility performance instead. Choosing a strategy means **prioritizing some dimensions and accepting limits on others** — there is no shape that touches every corner.

`[slides 3-9]`

## What is a supply chain, and who is in it?

A **Supply Chain (SC)** is *a flow of materials, information, and money across suppliers, manufacturers, and customers* — multiple entities, not one firm. **Supply Chain Management (SCM)** is managing that flow **efficiently and strategically**. The deck's basic picture runs Supplier → Factories → Warehouses → Outlets → Consumers, with materials, information and money moving in both directions along it.

Each stakeholder has a distinct role and a distinct interest:

| Stakeholder | Role | Key interest |
|---|---|---|
| **Suppliers** | Provide raw materials, components, or services | Stable demand, timely payment, long-term contracts |
| **Manufacturers/Producers** | Convert inputs into finished goods | Reliable supply, efficient operations, meeting schedules |
| **Distributors/Wholesalers** | Buy in bulk, sell smaller quantities to retailers | Fast turnover, reduced storage costs, steady demand |
| **Retailers** | Sell directly to consumers/businesses | Stock availability, competitive prices, satisfaction |
| **Logistics Providers** (e.g., DHL, FedEx, Maersk) | Transportation, warehousing, delivery | Efficient routing, timely delivery, cost optimization |
| **Customers/End Users** | Final recipients | Right product, quantity, time, condition |
| **Supporting Service Providers** (banks, ERP vendors, auditors) | Enable operations via finance, tech, consulting | Long-term partnerships, stable operations |

A **single company can hold more than one role at once** — Amazon is both a Retailer and a Logistics Provider; Tesla is both a Manufacturer and a Retailer (direct-to-consumer). A single company can also sit inside **several distinct supply chains simultaneously** — Apple runs separate iPhone, MacBook, and Accessories supply chains, gaining revenue diversification and shared resources, at the cost of having to balance different priorities and partners across each one.

`[slides 10-14]`

?check id=log-004

## From a linear chain to a network

The old, textbook view is a **one-way flow**: supplier → manufacturer → customer. The deck argues the reality is a **Supply Chain Network** — many-to-many connections, with multiple suppliers and customers at each tier, overlapping roles, and shared logistics, technology, and information flows. **CP Group (Thailand)** is the deck's network example: it is simultaneously a supplier (animal feed, seeds), a manufacturer (CP Foods), a retailer (7-Eleven Thailand, Lotus, Macro), and a logistics provider (cold-chain and transport). Because its retail arm sells both CP's own products and competitors', and it runs domestic and global chains at once, CP Group cannot be drawn as one straight line — it is a node in a network.

`[slides 15-16]`

## SCOR performance attributes and types of supply-chain strategy

The **SCOR (Supply Chain Operations Reference)** model gives five performance attributes for judging a whole network, and the wording matters because the terms are easy to confuse: **Reliability** is consistent, accurate, dependable outcomes; **Responsiveness (Speed)** is how fast the chain delivers; **Agility (Flexibility)** is the ability to adapt to changes in demand or supply; **Cost** is total cost to operate the chain; **Asset Management Efficiency** is efficient use of inventory, facilities, and capacity.

These attributes map onto named **types of SC strategy**, each fitted to a different product/demand profile:

| SC Strategy | Focus | Best for | Example |
|---|---|---|---|
| **Efficient** | Cost minimization | Commodities, stable demand | Walmart, Tesco |
| **Responsive** | Speed & flexibility | Fashion, collectibles, fresh food | Pop Mart |
| **Agile** | Customization & adaptability | Luxury goods, niche markets | Louis Vuitton, Dell |
| **Digital** | Data-driven optimization | E-commerce, tech-savvy firms | Amazon, JD.com |
| **Green (Sustainable)** | Environmental & ethical impact | Eco-conscious brands | Patagonia, IKEA |

The deck stresses **strategic fit**: business-process performance (micro, operational) has to drive supply-chain performance (macro, network), and the same trade-offs from Week 1 reappear at network scale — Cost ↔ Flexibility, Inventory (Cost) ↔ Speed, Reliability ↔ Cost.

`[slides 17-19]`

## Push vs. pull systems

**Push (make-to-stock):** each stage produces or ships based on a **forecast** — supplier supplies to forecast, manufacturer produces to forecast, distributor holds inventory based on forecast, retailer stocks based on forecast, and the customer simply buys what is available.

**Pull (make-to-order):** each stage reacts to **actual orders** — supply to order, produce to order, automatic replenishment triggered by real consumption, ending with the customer's own order.

| Feature | Push | Pull |
|---|---|---|
| Trigger | Forecast | Customer demand |
| Inventory | High | Low |
| Flexibility | Low | High |
| Risk | Overstock | Stockout |

Real chains usually combine both, separated by a **push-pull boundary**: the upstream **Push Segment** builds to forecast, and the downstream **Pull Segment** (closest to the consumer) reacts to real orders. The deck's case contrast is **Big C vs. Zara**: Big C is cost-leadership, predictable demand, and a **Push** model; Zara is differentiation (fast fashion), trend-driven demand, and a **Pull/Hybrid** model prioritizing speed and flexibility over pure cost efficiency.

`[slides 20-23]`

?check id=log-015

## Network design and sourcing strategy

**Supply Chain Network Design** decides the number and location of suppliers, factories, and warehouses, and whether the system is **centralized or decentralized**. The deck names four factors driving that choice: **cost, delivery speed, customer location, and risk management**.

Sourcing decisions carry their own trade-offs:

| Strategy | Description | Risk |
|---|---|---|
| **Single sourcing** | One trusted supplier | Vulnerable to disruptions |
| **Multiple sourcing** | Two or more suppliers | Higher coordination cost |
| **Global sourcing** | International cost advantage | Political/transport risk |
| **Local sourcing** | Nearby suppliers | Faster delivery, less risk |

**Global sourcing** specifically trades cost, quality, and capacity advantages against **lead time, disruption, and tariff** risk. The deck's named mitigation strategies are using **multiple suppliers** and **nearshoring** — moving sourcing closer to the market to cut that risk.

`[slides 24-26]`

## The bullwhip effect and the beer game

The **Bullwhip Effect**: small changes in real consumer demand cause **increasing variability upstream** as each tier over-reacts to the tier below it. The deck's named causes are **lack of visibility** into real demand, **poor forecasting**, **promotions**, and **batch ordering**. The results are overstock and waste at some tiers, stockouts and lost sales at others, and unhappy customers and suppliers throughout.

The class experiences this directly through the **Beer Game**: four roles per team (Retailer, Wholesaler, Distributor, Manufacturer) play 20 weekly rounds, each starting with 10 cases of inventory, receiving shipments, filling incoming orders, and placing a new order upstream — **without talking about demand forecasts** outside their own orders, which is exactly what starves the chain of visibility. Costs are **$0.50 per case of inventory per week** and **$1.00 per case of backlog (unmet demand) per week**, and the team with the lowest total cost wins. The debrief connects the swings in inventory and backlog the teams just experienced straight back to the causes above.

The deck's proposed fix is **real-time dashboards**: sharing live inventory, transit times, and sales data shrinks the visibility gap that drives the bullwhip effect and supports faster, data-driven decisions — illustrated by Zara's real-time store-to-HQ sales data and decentralized reordering, and Amazon's AI-based inventory prediction and live fulfillment dashboards.

> **Exam focus.** The bullwhip effect and the Beer Game are the deck's own case study: expect a scenario question naming a cause (e.g., batch ordering, no shared forecast) and a computation using the $0.50/$1.00 cost rule.

`[slides 27-33]`

?check id=log-019
