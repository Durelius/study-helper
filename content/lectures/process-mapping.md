---
id: process-mapping
title: "Operations Design & Process Mapping (BPMN)"
deck: "old 02"
slides: 33
summary: "How to design an operation along four dimensions, and how to draw its processes correctly in BPMN — the element families, the notation, and the modelling rules an examiner will ask you to spot a violation of."
---

## Operations and the four dimensions of operations design

**Operations** are how a business **transforms inputs into outputs**. An operation is made of **processes** (the sequence of steps), **resources** (people, equipment) and **capacity & flow** (how much, how fast). A barista making coffee, Toyota assembling cars and a hospital managing patient flow are all operations.

Designing an operation means answering four questions:

| Dimension | Question it answers |
|---|---|
| **Process** | What sequence of activities transforms inputs into outputs? |
| **Capacity** | How much can we produce in a given time? |
| **Flow** | How does work move through the system? (Fast, smooth, slow?) |
| **Bottlenecks** | Where do delays or constraints happen? |

`[slides 3-4]`

## The five common process types

| Process type | Example | Characteristics |
|---|---|---|
| **Project** | Construction | Unique, one-time |
| **Job shop** | Tailor shop | Custom, small batches |
| **Batch** | Bakery | Medium volume, repeated |
| **Flow / Line** | Assembly line | High volume, standardized |
| **Continuous** | Oil refinery | 24/7, high efficiency |

The list runs from low volume / high variety (project) to high volume / no variety (continuous). Starbucks sits near job shop–batch (each drink is made to order); McDonald's order-to-cash is closer to flow/line (standardised, assembled from pre-made components).

`[slides 5-7]`

## Why map a process, and what BPMN is

Process mapping **visualises how work is actually done**, helps identify **delays, waste and inefficiencies**, improves team communication, and enables better **project planning** and **customer experience**.

**BPMN** (Business Process Model and Notation) is an **OMG standard since 2005**, now **BPMN 2.0**. It is **flowchart-based** and serves **both conceptual and executable models**, supporting visualization, simulation (time, cost) and implementation. Tools named in the deck: Signavio, ARIS, Bizagi, **Camunda**, **bpmn.io** — and MS Visio / Excel / paper and pen.

`[slides 8-9]`

## The four core elements

BPMN has four core element families: **activity/task** (rounded rectangle), **event** (circle), **flow** (arrow) and **gateway** (diamond). Everything else — pools, lanes, artifacts — organises or annotates these four.

`[slide 10]`

## Activities

Activities **capture work (to-do)** performed in the process.

- A **task** is **atomic** — one indivisible unit of work, drawn as a plain rounded rectangle.
- A **sub-process** is **compound** — drawn as a rounded rectangle with a **+ marker**, and expands into a process of its own with its own start and end events.

Two rules attach to activities: **every activity must sit inside a lane**, and every activity must have a **unique name in "do something" form** (verb + object), e.g. *Prepare Order*, *Pack the Order*, *Obtain Payment*. The slide's rejected example puts a task called **"Manage Booking" in the Ticket Sales lane and a second task also called "Manage Booking" in the Hotel Sales lane** — the names are the right *shape* but they are not unique.

`[slide 11]`

?check id=bpm-027
?check id=bpm-028

## Events

An **event** is something that **"happens" (have-done)** during the course of a business process — a **trigger** or a **result** that **starts, interrupts or ends** the flow.

| | Start | Intermediate | End |
|---|---|---|---|
| Border | thin single line | double line | **thick** single line |
| Timer | circle with clock | double circle with clock | — |
| Message | **Message received** — open envelope | open or filled envelope | **Message sent** — filled envelope |

Two rules: **at least one start event and one end event for each process**, and a **unique event name in "something done" form** — *Order Received*, *Invoice handled*, *PO Fulfilled*, *Order Completed*. The rejected example on slide 13 is a pool containing only *Prepare Order → Deliver Order*: correct tasks, but **no start and no end event**.

`[slides 12-13]`

## Gateways

A **gateway controls how sequence flows interact as they converge and diverge** within a process. A gateway *represents control* — **if the flow does not need to be controlled, no gateway is needed**. Each type has a **split** (diverging) and a **join** (converging) form.

| Gateway | Symbol | Split behaviour | Join behaviour |
|---|---|---|---|
| **XOR (exclusive)** | diamond with **X** | Exclusive decision — **take one branch** | Exclusive merge — **proceed when one branch has completed** |
| **AND (parallel)** | diamond with **+** | Parallel split — **take all branches** | Parallel join — **proceed when all incoming branches have completed** |
| **OR (inclusive)** | diamond with **O** | Inclusive decision — **take one or several branches** | Inclusive merge — **proceed when all *active* incoming branches have completed** |

The deck's examples: XOR — *Check invoice for mismatches* splits into *Post invoice* / *Re-send invoice to customer* / *Block invoice*, all merged by an XOR join into *Park invoice*. AND — after *Proceed to security check*, *Pass security screening* and *Pass luggage screening* run in parallel and an AND join leads to *Proceed to departure level*. OR — *Check order line items* forwards a sub-order to the Amsterdam warehouse, the Hamburg warehouse, or **both**, and the OR join waits only for the branches that were actually taken.

Note where the **conditions** go: on the XOR/OR example they are written **as labels on the outgoing sequence flows** ("No mismatches", "order contains Amsterdam products"), never as activity boxes.

`[slides 14-18]`

> **Exam focus.** The named exam question is *"BPM — find one error, what is correct."* Gateway semantics and the three guidelines below are the highest-yield thing on this deck. Learn XOR/AND/OR split *and* join behaviour as six separate statements.

## Guidelines for gateways

1. **Use XOR or AND rather than OR.**
2. **Use gateways as a pair** — every split has a matching join of the same type.
3. **Only "1 in, multiple out" or "multiple in, 1 out" for each gateway — never "multiple in, multiple out".**

The rejected example ("Process 6") is exactly a violation of rule 3: a start event feeds an XOR split into *Task 3*, *Task 4* and *Task 5*, and all three run into a **single XOR gateway that then has two outgoing flows** to *Task 6* and *Task 7*. The fix is two gateways: a multiple-in/one-out join, then a one-in/multiple-out split.

`[slide 19]`

?check id=bpm-024
?check id=bpm-025

## Connecting objects

| Object | Line | Meaning |
|---|---|---|
| **Sequence flow** | solid line, filled arrowhead | Order of activities **within the same process or pool** |
| **Message flow** | dashed line, open circle at source, open arrowhead | Exchange of messages **between different processes or pools** |
| **Association** | dotted line | **Links artifacts** (text annotations, data objects) to flow objects **without affecting execution order** |

`[slide 20]`

## Pools and lanes

In BPMN a **swimlane** is divided into two types:

- **Pool** — a **participant / organisation** that takes part in the process. A rectangular container that holds flow objects (tasks, activities) laid out vertically or horizontally.
- **Lane** — a **function / department / sub-division / actor role within a pool**, used to organise and categorise the activities inside that pool.

A pool can be shown as a **white box**, with all details exposed, or as a **black box**, with all details hidden: an **empty box with no flow objects and no lanes**, whose **message flows link to its boundary**.

**Guidelines for swimlanes:** activities need to be **in a lane**; **sequence flows stay within a pool**; **message flows go across pools**.

The accepted example ("Process 7") is a food order: a *Process 7* pool with two black-box pools *Grab* and *Deliver*; a start event *Order Received* fed by a message flow "Order" from Grab; an XOR split labelled *Order includes Food* / *Order includes Beverage* into *Cook the Food* / *Make the Beverage*; an XOR join into *Pack the Order*; and an end event *Order Passed to Delivers* sending the message "Order Pack" to the Deliver pool. The rejected example ("Process 2") has an activity, *Write an Acceptance Letter*, **drawn straddling the line between the Manager and Staff lanes** instead of sitting inside one.

`[slides 21-23]`

?check id=bpm-031

## Worked example: the Order-to-Cash process

Two pools. **Customer** is a **black box**. **Order-to-Cash Process** is a white box with four lanes: **Sales, Production, Delivery, Finance**.

The flow: a **message start event "PO received"** in Sales (fed by the message flow "PO" from the Customer) → **Accept Order** (a sub-process; it sends "Request Information" to the Customer and receives "Informating" back) → **Produce** in the Production lane → **Deliver** in the Delivery lane, which sends the message "Products" to the Customer → **Obtain Payment** in the Finance lane, which receives the message "Payment" → end event **PO Fulfilled**.

Three things to notice: all four activities are **sub-processes** (+ marker), so the model is one level of a hierarchy; there is **no gateway at all**, because this flow never needs to be controlled; and every arrow crossing the pool boundary is a **dashed message flow**, while every arrow inside the pool — including the ones that cross lane boundaries — is a **solid sequence flow**.

`[slides 22, 24]`

## The "Any Errors?" exercise

A *Process 8* pool over a black-box *Customer* pool: start event *Order Received* → *Scan all products* → *Request Membership Code* (sending the message "Request Membership") → **XOR split with two unlabelled outgoing flows**, one to a box reading **"Customer has a Membership Code"** and one to a box reading **"Customer does not have a Membership Code"** → *Input the Membership Code* (receiving the message "Memership Code") → XOR join → *Request & Receive Payment* → end event *Order Fulfilled*.

The headline error: those two boxes are **conditions, not work**. They are drawn as activities but neither is a "do something" name. They belong as **condition labels on the gateway's outgoing sequence flows** — exactly as *Order includes Food* / *Order includes Beverage* were used on slide 23. A second, softer point: *Request & Receive Payment* bundles two actions into one task name.

`[slide 25]`

## Little's Law and process KPIs

**Little's Law:** **Inventory (WIP) = Throughput Rate (λ) × Flow Time (CT)**. At any time, the amount of work in the system depends on how fast things are flowing and how long they stay in the system.

| Term | Meaning | Example |
|---|---|---|
| **Inventory (WIP)** | Items/customers in process | 6 drinks being made |
| **Throughput** | Output per time | 3 drinks per minute |
| **Flow time** | Average time in system | 2 minutes from order to pickup |

**Process KPIs** help monitor, manage and improve performance:

| KPI | What it measures | Example |
|---|---|---|
| **Cycle time** | Time to complete one unit | 2 minutes/order |
| **Throughput rate** | Output per unit time | 30 drinks/hour |
| **Utilization** | % of resource used | Barista working 80% of shift |
| **Wait time** | Idle time for customer or item | 4 minutes queue time |
| **Defect rate** | % of incorrect outputs | 2 wrong orders/day |

> **Exam focus.** "BPM error **or quality KPI**" is one named exam question — be able to match each KPI to what it measures, and to say whether a long customer wait is caused by too much inventory, slow throughput or long flow time.

`[slides 28-31]`

## The modelling-rule checklist

| # | Rule | Slide |
|---|---|---|
| 1 | Activities must sit inside a lane | 11, 23 |
| 2 | Activity names are unique and in "do something" (verb + object) form | 11 |
| 3 | A task is atomic; a sub-process is compound and carries a + marker | 11 |
| 4 | At least one start event and one end event per process | 13 |
| 5 | Event names are unique and in "something done" form | 13 |
| 6 | A gateway is only needed where the flow must be controlled | 14 |
| 7 | Prefer XOR or AND over OR | 19 |
| 8 | Use gateways as a pair — every split has a matching join | 19 |
| 9 | Each gateway is "1 in, multiple out" or "multiple in, 1 out", never both | 19 |
| 10 | Branch conditions are labels on the outgoing sequence flows, not activity boxes | 15, 23, 25 |
| 11 | Sequence flows stay within one pool | 20, 23 |
| 12 | Message flows are used between pools | 20, 23 |
| 13 | Associations link artifacts only and carry no execution order | 20 |
| 14 | A black-box pool is empty — no flow objects, no lanes; message flows touch its boundary | 22 |
| 15 | A pool is a participant/organisation; a lane is a function/department/role within it | 21 |

`[slides 11-25]`
