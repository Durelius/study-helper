---
id: operations
title: "Operations Management"
deck: "new 02"
slides: 45
summary: "How value is actually created inside the firm: mapping a process, finding its bottleneck, measuring its flow and performance, and improving it."
---

## Where operations sits in the value chain

The course runs Value Chain → Operations → Supply Chain. Week 1 asks **where value is created**, Week 2 (this deck) asks **how value is created efficiently**, and Week 3 asks **how organisations coordinate flows externally**. Operations is one activity inside Porter's chain, but that single activity is performed through **many processes**, and each process is made of **activities (steps)**. The deck's worked hierarchy is: value-chain level (Inbound Logistics → Operations → Outbound Logistics → Marketing & Sales → Service) → process level (the Order-to-Coffee process: Receive Order → Process Payment → Prepare Drinks & Food → Deliver or Prepare Pickup) → activity level (Order → Pay → Prepare → Pickup).

The illustrative Thai case is **7-Eleven / CP ALL**: Inbound Logistics → Distribution Centres → Store Operations → Sales & Services → Customer Value, with support activities Procurement, Technology & Data, HR and the CP ALL network. The core value created is *right product, right place, right time, fast service*.

| Value chain | Process |
|---|---|
| High-level | Detailed |
| Whole organisation | Specific workflow |
| **Where** value is created | **How** value is created |
| Strategic | Operational |
| Example: Operations | Example: Order-to-Coffee process |

`[slides 3-5, 8]`

## Operations management and the transformation model

**Operations Management** = *designing, running, measuring and improving the processes that transform inputs into outputs*. All four verbs are examinable — a definition that says only "running the factory" is wrong.

The model is **Inputs → Transformation Process → Outputs**. Inputs are the resources used, the transformation is the process that adds value, outputs are the results created.

| Setting | Inputs | Transformation | Output |
|---|---|---|---|
| Starbucks | Coffee beans, labour, machine | Brewing, service, store experience | Coffee + customer experience |
| Hospital | Patient, staff, information | Diagnosis, treatment, care | Healthcare service |
| Food delivery | Order, rider, restaurant | Pick up, transport, handover | Delivered meal |

A **process** is *a sequence of activities that transforms inputs into outputs* — e.g. customer order → take order → prepare drink → serve → coffee.

`[slides 6-7]`

?check id=ops-001
?check id=ops-003

## Defining the process boundary: SIPOC

Before you can map a process you must say where it starts and stops. **SIPOC** = **S**upplier, **I**nput, **P**rocess, **O**utput, **C**ustomer, and its job is to *define what is inside and outside the process*. In the deck's diagram only **P (the process) is inside the boundary**; suppliers and inputs sit outside upstream, outputs and customers sit outside downstream.

For Starbucks: suppliers = coffee bean, dairy, syrup, cup, equipment and technology suppliers; inputs = beans, milk, syrups, cups, the customer order, payment, baristas, espresso machines, POS; process = Order → Prepare → Serve; outputs = prepared beverages, food items, receipts, customer experience, recorded transactions; customers = in-store, mobile-app and delivery-partner customers.

The pairing to remember: **SIPOC answers "what is the process?"; the process map answers "how does the process work?"** The deck's Starbucks order-to-coffee map has nine steps — Arrive, Queue, Order, Pay, Order sent to barista, Prepare, Quality check, Call/Notify, Pickup — split into **customer actions**, **frontstage Starbucks actions** and **backstage supporting actions**.

`[slides 9-11]`

## Process mapping and BPMN notation

**BPMN** = Business Process Model and Notation, an **OMG standard since 2005** (now BPMN 2.0), flowchart-based, usable for both **conceptual** and **executable** models; free tool: bpmn.io.

| Symbol | Meaning | Example |
|---|---|---|
| Oval | Start / End of a process | "Customer enters" … "Order completed" |
| Rectangle | Activity / Task | "Take Order" |
| Diamond | Decision | "Payment successful?" |
| Arrow | Flow / sequence | Order → Pay → Prepare → Pickup |
| Clock | Delay / waiting (BPMN timer event) | "Wait for drink (preparation time)" |

**Activities** capture the work (the to-do) performed in the process. Rules: every activity must sit **in a lane**, and every activity needs a **unique name phrased as "do something"** (verb + object). An activity is either a **Task** (atomic) or a **Sub-process** (compound, drawn with a **+** marker and expandable into its own flow).

**Events** are something that *happens* during the process — a trigger or a result that starts, interrupts or ends a flow. They come as **start, intermediate and end** events, and in message, timer and plain flavours. Rules: **at least one start and one end event per process**, and event names are phrased as **"something done"** (past participle), e.g. "Order received", "Invoice handled".

**Gateways** *control how sequence flows interact as they converge and diverge*. A gateway represents control only — **if the flow does not need to be controlled, no gateway is needed**. Each gateway type appears as a **split** and a matching **join**.

| Gateway | Symbol | Behaviour | Deck example |
|---|---|---|---|
| **XOR** (exclusive) | Diamond with × | Exactly **one** outgoing path is taken | Invoice check: no mismatches → post invoice; correctable → re-send to customer; not correctable → block invoice |
| **AND** (parallel) | Diamond with + | **All** outgoing paths run; the join waits for all | Airport: pass security screening **and** pass luggage screening, then proceed to departure level |
| **OR** (inclusive) | Diamond with ○ | **One or more** paths, depending on conditions | Order lines forwarded to the Amsterdam warehouse, the Hamburg warehouse, or both |

> **Exam focus.** "BPM error" is one of the five named exam questions. The three errors this deck plants explicitly are: an activity **not placed in a lane**, a process with **no start and/or end event**, and **mis-named** elements (an activity named like an event, or a non-unique name). Adding a gateway where the flow needs no control is the fourth thing to watch for.

`[slides 12-20]`

?check id=ops-011
?check id=ops-016

## A process contains more than work

Mapping reveals that only part of a process is actual work. The deck splits the coffee process into four kinds of step:

| Step type | What it is | Example | Adds value? |
|---|---|---|---|
| **Waiting** | The customer or process is idle | Queuing during busy hours | No direct value, but affects satisfaction and time |
| **Processing** | Actual work that transforms input into output | Preparing the coffee | **Yes** — this is what the customer pays for |
| **Decision** | A choice that determines the next step | Hot or iced? | **Yes** — ensures the right product is made |
| **Movement** | People, information or items moved from place to place | Order transferred to the preparation station | No direct value; necessary but adds time and cost |

This feeds the three-way classification used in improvement work:

- **Value-adding** — changes the product/service in a way the customer values. *Brewing coffee.* ("This is why I came.")
- **Necessary but non-value-adding** — required for the process to work, but creates no direct customer value. *Payment processing.* ("I don't mind doing this, but it's not why I came.")
- **Waste** — consumes resources without creating value. *Waiting for an unavailable cup.* ("Frustrating and adds no value to me.")

The trap: payment processing is **not** waste. It is necessary. Waste is the step you could delete and lose nothing but cost and delay.

`[slides 21-22]`

## Capacity, bottlenecks and utilisation

**Capacity** = *the maximum output a process or resource can produce in a given period*. It is always **quantity per unit of time**: 25 drinks/hour, 400 units/day, 50 patients/day, 2,000 transactions/second. The deck's explicit teaching point: **"100 units" is not capacity; "100 units/hour" is capacity.**

The deck's running example gives four serial steps: Order 40/hr, Payment 50/hr, Preparation 25/hr, Pickup 60/hr.

> **Process Capacity = min(40, 50, 25, 60) = 25 customers/hour.**

The capacity of a **serial** process is limited by its **lowest-capacity step**. That step is the **bottleneck**: *the step that constrains the throughput of the overall process*. Preparation is the bottleneck here.

The deck's key thought experiment: if payment capacity is doubled from 50 to 100 customers/hour, **process capacity does not change** — it is still 25/hour, because you invested in a non-bottleneck. Only adding capacity **at the bottleneck** raises output, and when you do, the bottleneck moves to whichever step is now lowest.

**Utilisation = Demand ÷ Capacity × 100%.** With capacity held at 50 customers/hour:

| Demand | Utilisation | What happens |
|---|---|---|
| 30/hr | 30/50 = **60%** | Unused capacity, no waiting; shorter wait times, lower stress, higher flexibility — but higher cost per output because resources are idle |
| 48/hr | 48/50 = **96%** | Demand ≈ capacity; efficient operation, very few waiting, but little buffer and **sensitive to variation or disruption** |
| 60/hr | 60/50 = **120%** | Demand exceeds capacity; queues, delay, frustration and **possible lost sales** |

> **Exam focus.** "Is 100% utilisation always desirable?" is asked directly on a slide. The answer the deck sets up is no: high utilisation is efficient but leaves no buffer, so any variation turns immediately into queues and lost sales.

`[slides 23-29]`

?check id=ops-022

## Flow: throughput, flow time, WIP and Little's Law

Three measures describe how work flows:

| Measure | Question it answers | Unit |
|---|---|---|
| **Throughput rate** | How many units are completed per unit of time? | units / time |
| **Flow time** | How long does one unit spend in the process? | time / unit |
| **Work-in-process (WIP)** | How many units are currently in the process? | units |

**Little's Law** states that in a **stable** process:

> **WIP = Throughput Rate × Flow Time**

Worked example from the deck: a packing line with throughput rate 30 units/hour and flow time 2 hours has WIP = 30 × 2 = **60 units** — on average 60 units are in the system at any point in time. Rearranged, Flow Time = WIP ÷ Throughput and Throughput = WIP ÷ Flow Time, which is how you diagnose a long wait: is the queue long (high WIP), is the process slow (low throughput), or does each unit linger (long flow time)?

`[slides 30-32]`

## Performance objectives and KPIs

"Good operations" is measured on five objectives:

| Objective | Question | What it looks like |
|---|---|---|
| **Quality** | Do we do it correctly? | Meet requirements, low defects, right first time, customer satisfaction |
| **Speed** | Do we do it quickly? | Short cycle time, fast response, high throughput, less waiting |
| **Dependability** | Do we do it when promised? | On-time delivery, consistent performance, reliable processes, trust |
| **Flexibility** | Can we adapt? | Handle variety, adjust to change, scale up or down, support new needs |
| **Cost / Productivity** | Do we use resources efficiently? | Lower cost, higher output per input, eliminate waste, better resource use |

KPIs are derived, not invented. The chain is **Business/Operational Objective → Process Driver → KPI → Target**: "Serve customers faster" → "Reduce preparation delay" → "Average preparation time" → "< 4 minutes".

| Performance objective | KPI |
|---|---|
| Speed | Cycle time |
| Quality | Defect rate |
| Dependability | On-time completion |
| Productivity | Units / labour-hour |
| Capacity | Utilisation |
| Reliability | Downtime |
| Customer flow | Waiting time |

> **Exam focus.** "Quality KPI" is part of one of the five named exam questions. Be able to pair an objective with its KPI both ways, and remember that a KPI is *a signal to understand, decide and improve*, not just a number: Metric ("waiting time = 12 min") → **So what?** → Insight ("cause: preparation bottleneck") → **Now what?** → Decision ("add barista, redesign station layout, simplify menu") → **So that** → Outcome ("waiting time ↓, customers happier, sales ↑").

`[slides 33-36]`

?check id=ops-030

## Continuous improvement

**Continuous improvement** = *the systematic effort to improve process performance over time*: Current Process → Measure → Identify Gap → Improve → New Process.

**PDCA** is the deck's simple improvement cycle:

- **Plan** — identify the problem and plan the solution: understand the current process, set objectives, analyse root causes, plan improvements.
- **Do** — implement the plan **on a small scale**: execute, train and communicate, collect data, document what happened.
- **Check** — monitor results and evaluate performance: measure results, compare with objectives, identify gaps, determine root causes.
- **Act** — standardise what works and act on what doesn't: standardise improvements, update procedures, share lessons learned, start the cycle again.

Two root-cause tools, and the difference between them is examinable:

- **5 Whys** — *drill down one causal chain*. Problem: customers wait 12 minutes → Why? coffee preparation is slow → Why? one machine becomes overloaded → Why? most popular drinks use the same machine → Why? the menu and workflow are not balanced → Why? **the system was designed for average demand, not peak demand** (the root cause).
- **Fishbone / Ishikawa** — *explore multiple possible cause categories* in breadth: **People, Machine, Method, Materials, Measurement, Environment**, all pointing at the effect ("customers wait 12 minutes for coffee").

The full **Process Improvement Cycle** is five stages: **1 Map** (define start and end, list steps in order, identify handoffs) → **2 Analyze** (find bottlenecks, spot waste such as motion, waiting and rework, identify root causes) → **3 Measure** (choose key metrics, collect baseline data, understand current performance) → **4 Improve** (generate solutions, prioritise and select, implement changes) → **5 Measure Again** (measure new results, compare to baseline, confirm improvement). Repeat.

When judging any proposed improvement, the deck gives three test questions: **Does it address the bottleneck / root cause? Which KPI should change? Could it create a new problem elsewhere?**

`[slides 37-45]`

?check id=ops-034
