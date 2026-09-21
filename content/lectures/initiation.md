---
id: initiation
title: "Project Initiation & Scope Management"
deck: "new 05"
slides: 39
summary: "How an approved project becomes a defined project: charter, stakeholders, requirements, scope statement, WBS and RACI."
---

## Where this lecture sits

Week 4 answered **why should we do the project?** (selection → business case → approval).
This week answers **what exactly will we deliver, and who needs to be involved?**
Week 6 answers **when, with what resources, and how much?**

After approval the project must still be *defined*, in four moves: **authorize** (who has
authority?), **understand** (who needs what?), **define** (what is included?), and
**decompose** (what work must be done?). The slide's one-line summary is the point:
**approval gives permission, scope gives clarity.**

`[slides 3-5]`

## The project charter

**Project charter**: a document that *formally authorizes the project* and *gives the
project manager authority to lead it*. The deck reduces it to three verbs:

| Verb | What it achieves |
|---|---|
| **Authorize** | The project may begin |
| **Align** | Sponsor and PM share the same high-level understanding |
| **Empower** | The PM receives authority to coordinate the work |

The summary line on the slide is blunt: **Charter = AUTHORIZATION**.

Seven things go in a charter:

1. **Project purpose / business need** — why?
2. **Objective** — what outcome?
3. **High-level scope** — what are we broadly creating?
4. **Major deliverables** — what will exist?
5. **Sponsor & project manager** — who owns / who leads?
6. **High-level milestones & budget** — what are the boundaries?
7. **Major risks / assumptions** — what should management know?

The **FreshMart charter example**: business need = inventory accuracy only 88%, stockouts
harm omnichannel reliability; objective = raise inventory accuracy to 97% with real-time
stock visibility; high-level deliverable = real-time inventory-management capability across
25 stores; sponsor = the **COO**; PM = assigned PM; target = 9 months; budget = ฿6 million;
major risks = ERP integration, staff adoption, data quality.

> **Exam focus.** "Project definition & planning" almost certainly starts at the charter.
> Know that the *sponsor* authorizes and signs it, the *PM* is named and empowered by it,
> and that it is short and high level.

`[slides 6, 8, 9, 11]`

## Charter vs business case, charter vs plan

Two separations the deck makes explicitly, and both are classic trap questions.

| | Business case | Project charter |
|---|---|---|
| Question answered | **WHY** should we invest? | **WHAT** are we authorizing? |
| Timing | *Before* approval | *After* approval / at initiation |
| Focus | Need, benefits, cost, risk, feasibility | Objective, high-level scope, sponsor, PM, key milestones, major risks, high-level budget |
| Output | Decision to invest | Authority to begin |

The charter is **not** the detailed plan. The charter holds *why, what, who, major
boundaries*. Detailed requirements, the WBS, the schedule, resources, budget and risk
responses all come **later**. The slide's rule: **authorize first, detail later.**

`[slides 7, 10]`

?check id=ini-003

## Stakeholders

**Stakeholder**: a person, group, or organization that *can affect, be affected by, or
perceive itself to be affected by* the project. The two questions to ask are **"who
cares?"** and **"what do they need?"**

Typical stakeholder groups around a project: **sponsor, users, customers, project team,
operations, IT, suppliers**. The deck's governing claim: **scope cannot be defined
correctly without understanding stakeholders** — they are the ones who define "success",
and they will not all want exactly the same thing.

FreshMart's stakeholder map: sponsor = COO; users = store employees / inventory planners;
customers = online and in-store customers; operations = store managers / distribution;
technical = IT / ERP team; external = the technology vendor.

`[slides 12-14]`

## From stakeholder need to requirement

The deck gives one chain, and the exam can test any link of it:

**Stakeholder** (store manager) → **Need** (know real stock levels) → **Requirement**
(inventory updates within X minutes) → **Deliverable / feature** (real-time inventory
dashboard).

A need is vague and human; a **requirement** is stated precisely enough to be built and
tested. Three requirement types:

| Type | Question | FreshMart example |
|---|---|---|
| **Functional** | What must it *do*? | Update store inventory after every transaction |
| **Performance / quality** | How well must it perform? | Inventory accuracy ≥ 97% |
| **Constraint** | What must it *comply with*? | Integrate with the existing ERP |

`[slides 15-16]`

## Project scope vs product scope

**Project scope**: the work required to create the agreed project deliverables. It answers
three questions — what are we delivering, what is included, what is excluded — and a good
scope creates a **shared boundary**.

| | **Product scope** | **Project scope** |
|---|---|---|
| Keyword | **WHAT** | **WORK** |
| Meaning | Features and functions of the result | Work required to create that result |
| FreshMart example | Real-time inventory dashboard, ERP synchronization, stock alerts | Requirements, design, integration, testing, training, deployment |

**Product scope describes the result. Project scope describes the work.**

The chain from objective downward: **objective** (accuracy to 97%) → **deliverables**
(inventory platform, ERP integration, dashboard, training) → **scope** (the work required
to create and deploy those deliverables) → **work packages** (manageable pieces of work).

`[slides 17-19]`

## The scope statement and acceptance criteria

**Scope statement = WHAT + BOUNDARIES.** Six components:

| Component | Question |
|---|---|
| **Objective** | What outcome? |
| **Deliverables** | What must be created? |
| **Requirements** | What must the output do? |
| **In scope** | What work is included? |
| **Out of scope** | What is specifically excluded? |
| **Acceptance criteria** | How will we know it is acceptable? |

FreshMart's scope statement: objective = raise inventory accuracy from 88% to ≥97%;
deliverables = real-time inventory database, ERP interface, store inventory dashboard, user
training; **in scope** = 25 stores, inventory synchronization, stock alerts, training;
**out of scope** = replacing the ERP, a new warehouse, customer mobile-app redesign;
acceptance criteria = accuracy ≥97%, approved integration test, all stores deployed.

**Acceptance criteria** answer "how do we know the deliverable is done?" — for FreshMart:
inventory accuracy ≥ 97%, system response within the agreed response time, ERP
synchronization passes testing, 25 stores operating successfully. **A deliverable without
acceptance criteria is difficult to control.**

**Deliverables ≠ activities.** Deliverables are **outputs** (what exists): inventory
dashboard, ERP interface, training package. Activities are **work** (what we do): design
dashboard, develop interface, conduct training.

`[slides 20-23]`

## Scope creep, gold plating and controlled change

**Scope creep**: the *uncontrolled* expansion of project scope **without a corresponding
adjustment to time, cost, or resources**. A feature here, a requirement there, a "small
request", an extra report — and scope rises, which drives **time ↑, cost ↑, risk ↑**.

Five causes named in the deck:

1. **Unclear requirements** — vague, incomplete or frequently changing.
2. **Users involved too late** — key users are not engaged early and add requests later.
3. **Complexity underestimated** — the work is harder than expected.
4. **Weak change control** — changes accepted without review or impact assessment.
5. **Gold plating** — adding extra features or enhancements that are **not necessary**.

Note the distinction the examiner will exploit: **gold plating is extras the team adds on
its own initiative**; scope creep is the broader uncontrolled growth, often driven by
stakeholder requests.

**Scope change ≠ scope creep.** Good project management does not prevent change — it
*controls* change.

| **Controlled change** | **Scope creep** |
|---|---|
| Request → **impact analysis** → **decision** → **update the plan** | Request → "just add it" → work changes → **no adjustment to time / cost** |

Worked example: management asks "can we add AI demand forecasting to the inventory
project?" The correct response is to assess four impacts — **scope** (new feature not in
current scope, needs extra data sources, may affect other requirements), **schedule**
(extra analysis, model development and testing; likely delay to go-live), **cost** (extra
software/AI tools, data preparation, training and support), **risk** (model accuracy, data
quality, technical/integration complexity, resource overload) — and then **approve**
(add now), **defer** (do it in a future phase), or **reject** (do not add to this project).

> **Exam focus.** For the case study, the expected answer to any "can we just add…" prompt
> is: assess scope/schedule/cost/risk impact, then approve, defer or reject — never "just
> add it", and never simply refuse.

`[slides 24-27]`

?check id=ini-019

## The Work Breakdown Structure

"Build the inventory system" is too large to manage, so you **break it down**. **You cannot
schedule, assign, or estimate work that has not been defined.**

**WBS**: a **hierarchical decomposition of the project scope into smaller, manageable
components**. The levels run **project → deliverables → sub-deliverables → work packages**,
and **the lowest manageable level is the work package**.

**The 100% rule — WBS = 100% of the project scope:**

- **Includes all required work** — every deliverable, activity and task needed to complete
  the project is included.
- **No missing work** — all work required to meet the project objectives is captured.
- **No duplicated work** — each piece of work appears **only once** in the WBS.

And the trap: **the WBS does NOT show sequence.** It is a breakdown, not a schedule; order
and dependencies come later, in Week 6.

**PBS vs WBS**: a **Product** Breakdown Structure decomposes the **product** ("what are we
building?" — Inventory System → dashboard, database, ERP interface, alerts). A **Work**
Breakdown Structure decomposes the **project work** ("what work is required?" — ERP
Interface → define integration requirements, build API, test synchronization, deploy
interface).

FreshMart WBS (Level 1 = project, Level 2 = major deliverables, Level 3 = work packages):

| Level 2 | Level 3 work packages |
|---|---|
| 1.1 Requirements & Design | Requirements approved; Architecture designed |
| 1.2 Inventory Platform | Database; Dashboard; Alerts |
| 1.3 ERP Integration | Interface; Data mapping; Testing |
| 1.4 Deployment | Pilot; Training; Rollout |
| 1.5 Project Management | Coordination; Status reporting; Closure |

**When do we stop decomposing?** When a work package is small enough that we can:
**estimate it** (reasonably estimate time and cost), **assign it** (to *one* accountable
owner), **monitor it** (track and measure progress during execution), and **accept it**
(a clear, objective way to know when it is done). Decompose until the work is manageable —
**not until every tiny action appears.**

`[slides 28-33]`

?check id=ini-026

## From WBS to responsibility: RACI

The **responsibility matrix** joins the **WBS** ("what work?") to the **OBS**, the
organizational breakdown structure ("who is available?"), and answers "who owns each work
package?" — **WHAT + WHO**.

**RACI** letters:

| Letter | Role | Meaning |
|---|---|---|
| **R** | Responsible | **Does the work** |
| **A** | Accountable | **Owns the result** |
| **C** | Consulted | **Provides input** |
| **I** | Informed | **Needs updates** |

The rule stated on the slide: **one clear Accountable owner per work item.** Several people
may be *Responsible* for the same work package, but exactly one is *Accountable*.

FreshMart RACI:

| Work package | Sponsor | PM | IT | Store Ops | Vendor |
|---|---|---|---|---|---|
| Requirements | I | **A** | R | C | C |
| ERP Interface | I | **A** | R | C | R |
| Pilot | I | **A** | C | R | C |
| Training | I | **A** | C | R | R |
| Approval | **A** | R | C | C | I |

Read two things off it: ERP Interface has **two** Responsible parties (IT and the vendor)
but still only one Accountable; and on the Approval row the roles flip — the **sponsor** is
Accountable and the **PM** is merely Responsible.

> **Exam focus.** Expect a scenario asking you to assign letters. The safe moves: the doer
> gets R, the single owner who signs off gets A, anyone whose expertise you must ask gets
> C, anyone who only needs to hear about it gets I — and never two A's in one row.

`[slides 34-36]`

?check id=ini-031

## Why scope comes first

**Scope** (what work?) → **WBS** (work packages) → and only then **schedule** (when?),
**resources** (who/what?) and **cost** (how much?). **Schedule and cost are built from
defined work — not guessed independently.**

Wrap-up in the deck's own words: the charter authorizes the project; stakeholders define
needs; requirements clarify expectations; scope defines boundaries; the WBS breaks down the
work; RACI assigns responsibility; and defined scope enables schedule and cost planning.

`[slides 37, 39]`
