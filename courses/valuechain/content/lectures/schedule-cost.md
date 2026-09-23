---
id: schedule-cost
title: "Project Planning: Schedule & Cost"
deck: "new 06"
slides: 47
summary: "How an approved scope becomes a sequenced schedule, a resourced plan and an approved budget — and how to find the critical path."
---

## The planning logic

Planning converts scope into an **executable commitment**. The deck runs it as eight numbered
steps grouped under four questions: **scope** (what will be delivered?), **schedule** (when will
it be done?), **resources** (with what?) and **cost** (how much?).

| # | Step | What it does |
|---|---|---|
| 1 | **WBS** | defines the work |
| 2 | **Activities** | make the work schedulable |
| 3 | **Dependencies** | determine sequence |
| 4 | **Duration** | determines timing |
| 5 | **Critical path** | determines project duration |
| 6 | **Resources** | make the schedule feasible |
| 7 | **Cost** | quantifies the required money |
| 8 | **Baseline** | creates the commitment |

A **WBS is not a schedule**: the WBS is hierarchical and carries no sequence, while the schedule
is time-based and shows sequence and dependencies. `[slides 5-8, 47]`

## Activities, work packages and deliverables

An **activity** is a defined piece of project work that consumes time and usually resources.
The deck's three-level worked example is worth memorising because it is an easy exam trap:

| Level | Example |
|---|---|
| **Deliverable** | Mobile application |
| **Work package** | Payment module |
| **Activity** | Integrate payment gateway |

A **milestone** is different again: a significant project event or checkpoint with **zero
duration**. It marks a moment (e.g. "Project approval"), it does not consume time. `[slides 14, 23]`

?check id=sc-003

## The four dependency types

A **dependency** describes how one activity affects another activity. All four types are named
from the predecessor's event to the successor's event.

| Type | Rule | Deck's event example |
|---|---|---|
| **Finish-to-Start (FS)** | Task B cannot **start** until Task A **finishes** | Rigging & Lighting → Decor |
| **Start-to-Start (SS)** | Task B cannot **start** until Task A **starts** | Catering Prep → Guest Arrival |
| **Finish-to-Finish (FF)** | Task B cannot **finish** until Task A **finishes** | Live Band Performance → Audio Recording |
| **Start-to-Finish (SF)** | Task B cannot **finish** until Task A **starts** | Day Shift Security → Night Shift Security |

FS is the default, and the deck bolds it. Read the name literally and you cannot get these wrong:
the *first* word is the predecessor's event, the *second* word is the successor's event.

> **Exam focus.** Classification questions give you a one-line scenario and ask for the type.
> The reliable method is to ask which *event* of each activity is constrained — a start or a
> finish — never to guess from whether the tasks look "parallel".

`[slides 15-17]`

?check id=sc-009

## Leads, lags and parallel work

- **Lead**: the successor starts **before** the predecessor completely finishes. Deck example:
  everyone else joins the dance floor before the bride-and-father dance ends.
- **Lag**: a deliberate **wait** between activities. Deck example: serve dinner, then prepare
  dessert after a wait.

A lead compresses the schedule; a lag stretches it. The deck also shows that moving work from
**sequential** (A → B → C, 10 days) to **parallel** (A → B and A → C, 7 days) shortens the
project, because only the longer of the two parallel branches adds to A. `[slides 18-20]`

## Effort vs duration

**Effort** is how much work there is (e.g. **40 person-hours**). **Duration** is how long it
takes on the calendar (e.g. **5 working days**). They are not the same number and not the same
unit. Crucially, **adding people does not always reduce duration proportionally** — two people
do not cut duration exactly in half, because of overheads, dependencies and coordination time.

Duration estimates come from **work quantity**, **resource availability**, **experience**,
**complexity** and **uncertainty**. `[slides 21-22]`

## Gantt chart vs network diagram

A **Gantt chart** is a visual bar chart of activities against time; it shows the activity, its
start or finish, duration, lead or lag, and milestones. A **network diagram** is a flowchart of
activities and dependencies. The deck's one-line summary is the examinable sentence:

> **Gantt = communicate time. Network = analyze logic.**

`[slides 24-25]`

## The AON box and the CPM method

In an **Activity-On-Node (AON)** diagram every activity is a three-row box. **Learn the layout —
the exam gives you a blank grid in exactly this shape.**

| Row | Left cell | Middle cell | Right cell |
|---|---|---|---|
| Top | **ES** — Earliest Start | **Duration** | **EF** — Earliest Finish |
| Middle | — | **Activity label** (spans the row) | — |
| Bottom | **LS** — Latest Start | **Float** | **LF** — Latest Finish |

- **EF = ES + Duration**
- **LF = LS + Duration**
- **Total float = LS − ES = LF − EF**

The **Critical Path Method (CPM)** has four steps, in this order:

1. **Forward pass** (earliest times). Start at 0, move left to right, `EF = ES + Duration`. Where
   several arrows **merge** into one activity, its ES is the **maximum** EF of its predecessors —
   it cannot start until *every* predecessor is done.
2. **Backward pass** (latest times). Start from the project finish, move right to left,
   `LS = LF − Duration`. Where an activity feeds several successors, its LF is the **minimum**
   LS of those successors — it must be finished in time for the *tightest* one.
3. **Float calculation**: `LS − ES` (equivalently `LF − EF`).
4. **Critical path identification**: the chain of activities with **float = 0**.

**Float (slack)** is the amount of time an activity can be delayed without delaying project
completion. The **critical path** is the **longest** sequence of dependent tasks, and any delay
on it delays the whole project.

> **Exam focus.** The first exam question is a critical path problem. Do the forward pass for
> *every* activity before starting the backward pass, and remember the two merge rules —
> **max on the way forward, min on the way back**. The single most common error is taking the
> minimum at a forward-pass merge.

`[slides 26-29]`

## Worked example: the deck's CPM network (slide 30)

The network is: **START** → A (6), B (4), F (10); **A** → C (3); **B** → D (4) and E (3);
**C** → H (2); **D** → H; **E** → G (3); **F** → G; **H** and **G** → **FINISH**.

**Step 1 — forward pass.** A, B and F all start at 0, so EF(A) = 6, EF(B) = 4, EF(F) = 10.
Then C starts at 6 → EF 9; D starts at 4 → EF 8; E starts at 4 → EF 7. H merges C and D, so
ES(H) = max(9, 8) = **9** → EF 11. G merges E and F, so ES(G) = max(7, 10) = **10** → EF 13.
Project duration = max(11, 13) = **13 days**.

**Step 2 — backward pass.** Set LF = 13 at the finish. LF(H) = 13 → LS 11; LF(G) = 13 → LS 10.
LF(C) = LS(H) = 11 → LS 8; LF(D) = 11 → LS 7; LF(E) = LS(G) = 10 → LS 7; LF(F) = 10 → LS 0.
LF(A) = LS(C) = 8 → LS 2. B feeds both D and E, so LF(B) = min(7, 7) = 7 → LS 3.

**Step 3 — float.** Float = LS − ES (check it against LF − EF) for every activity:

| Activity | ES | Dur | EF | LS | Float | LF | Status |
|---|---|---|---|---|---|---|---|
| A | 0 | 6 | 6 | 2 | 2 | 8 | |
| B | 0 | 4 | 4 | 3 | 3 | 7 | |
| C | 6 | 3 | 9 | 8 | 2 | 11 | |
| D | 4 | 4 | 8 | 7 | 3 | 11 | |
| E | 4 | 3 | 7 | 7 | 3 | 10 | |
| F | 0 | 10 | 10 | 0 | **0** | 10 | **critical** |
| G | 10 | 3 | 13 | 10 | **0** | 13 | **critical** |
| H | 9 | 2 | 11 | 11 | 2 | 13 | |

**Step 4 — the critical path is START → F → G → FINISH, 13 days.** Check it against the path
totals: A–C–H = 11, B–D–H = 10, B–E–G = 10, F–G = **13**. Note that the critical path here has
only **two** activities — longest means longest in *time*, not in number of boxes.

Because A has 2 days of float, a 2-day overrun on A costs nothing; a 3-day overrun pushes
A–C–H to 14 and the project finishes **1 day late**. That is what float buys you, and its limit.

The **project manager's attention** follows float: **highest attention to critical activities**,
**some flexibility for near-critical activities**, **greater flexibility where float is larger**.
`[slides 30-31]`

?check id=sc-032

## Resource planning

**Resource planning** determines **what** resources are required, **when** they are needed and
**whether** they are available, in five steps: activity → resource requirement → availability →
assignment → feasible schedule. The four resource types are **people, equipment, materials and
services**.

Two failure modes: **requirement > availability** (3 developers needed, 2 available), and
**resource over-allocation** (a developer with 8 h/day capacity assigned 6 h to Project A and
5 h to Project B = 11 h, an over-allocation of **3 hours**). **Resource leveling** fixes them by
reassigning work, delaying an activity, changing sequence, adding capacity or outsourcing.
`[slides 32-37]`

## Cost planning

**Project cost planning** estimates the financial resources required and establishes the approved
budget, in five steps: **estimate** activity costs → **aggregate** them → **add reserves** →
**approve** → **baseline**.

**Direct costs** can be traced directly to the project (labour, materials, equipment,
subcontractors); **indirect costs** support the project but are not tied to a single deliverable
(administration, office support, shared IT, overhead). *Direct costs build the deliverables;
indirect costs keep the organisation running.*

| Method | How it works |
|---|---|
| **Analogous** | Use historical data from similar past projects |
| **Parametric** | **Rate × quantity** (e.g. cost per square foot, per kitchen) |
| **Bottom-up** | Estimate each item or task and add them together |
| **Three-point** | `CE = (Cₒ + 4C𝑀 + C𝑃) / 6` from optimistic, most likely and pessimistic |

The deck's worked estimate: Design 40 h × ฿1,500 = ฿60,000; Develop 100 h × ฿2,000 = ฿200,000;
Test 30 h × ฿1,000 = ฿30,000 — **total ฿290,000**. `[slides 38-42]`

## Budget, baselines and the cost baseline

**Estimate = expected cost. Budget = approved financial commitment.** The build-up runs:
activity cost estimates + activity contingency reserve → work package cost estimates
+ contingency reserve → **control accounts** → **cost baseline** + **management reserve** →
**project budget**. The management reserve sits *outside* the cost baseline.

A **baseline** is the approved version of the project plan used as the reference for later
control. There are three: **scope baseline** (what?), **schedule baseline** (when?) and
**cost baseline** (how much?). The **cost baseline** shows how planned expenditure accumulates
over time — the familiar S-curve of cumulative cost against time, against which actual
expenditure is later compared.

> **Note.** This deck stops at the baseline. It does **not** cover earned value
> (PV / EV / AC / CV / SV / CPI / SPI / EAC) — those formulas are not examinable from new 06.

A good project plan answers four questions: **WHAT** must be done, **WHEN** must it happen,
**WHO/WHAT** resources are needed, and **HOW MUCH** will it cost. `[slides 43-47]`
