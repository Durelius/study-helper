---
id: risk-comm-quality
title: "Risk, Communication & Quality"
deck: "new 07"
slides: 52
summary: "How a project plan protects itself: naming uncertainty correctly, routing information to the right stakeholder, and defining in advance what 'acceptable' means."
---

## Where risk, communication and quality sit

Scope, schedule and cost say *what* the project will do. They do not say what happens when reality disagrees. The three remaining planning components are all produced in the **Planning** phase — **risk plan**, **communication plan**, **quality plan** — and are then used in Execution (manage team communication, perform quality assurance, implement risk responses) and Monitoring & Control (monitor quality, monitor risks, track communication).

Each answers one question: **Risk — what could happen?** **Communication — who needs to know what?** **Quality — what does acceptable mean?**

`[slides 2-4]`

## What a project risk is — and five things that are not risks

A **project risk** is an **uncertain** event or condition that may **affect** one or more project objectives. It can be **positive** (new technology, cost saving) or **negative** (supply delay, scope creep), and a single risk may have **one or more causes** and **one or more impacts**. Risk can strike any knowledge area: scope, time, cost, quality, communication, procurement, human resources, integration.

The exam's favourite trap is mixing risk up with its neighbours. Learn this table by its **timing** column — the timing alone usually settles the answer.

| Term | Definition | Timing | Example |
|---|---|---|---|
| **Risk** | An uncertain event that may affect the project (positive or negative) | **Future (potential)** | Possible delay due to supplier shipment problems |
| **Issue** | A problem that is currently affecting the project | **Present (actual)** | A supplier has missed a delivery deadline |
| **Constraint** | A limiting factor that restricts project options or decisions | **Always present** | Project must be completed within 3 months |
| **Assumption** | Something believed to be true without proof, used for planning | **Before or during** | Assuming all team members are available full-time |
| **Dependency** | A task or event that relies on another to start, continue or finish | **Ongoing relationship** | Task B starts only after Task A is completed |
| **Defect** | A flaw or bug in a deliverable that fails to meet quality standards | **After creation** | The login feature in the app doesn't work as intended |

Two pairings are worth rehearsing. A **risk becomes an issue** the moment it actually happens. An **assumption** and a **risk** come as a couple: the belief is the assumption, the possibility that the belief is wrong is the risk.

> **Exam focus.** One of the five named exam questions is "here is a case — what type of risk is this?". Half of those cases are not risks at all: they are issues, constraints, assumptions, dependencies or defects. Read the tense of the sentence first. Already happening → issue. A flaw inside something already built → defect. A fixed limit → constraint. A belief used for planning → assumption. A "cannot start until" link → dependency.

`[slides 5-6]`

?check id=rcq-004
?check id=rcq-009

## Types of project risk by source, and overarching risk

Once something *is* a risk, classify it by **source**. The deck gives five categories with fixed examples — memorise which example sits under which heading, because several are counter-intuitive.

| Source | Examples given |
|---|---|
| **Financial** | Budget overruns; resource misallocation |
| **Technical** | Technology failure or underperformance; unsuitable technology choices; integration issues |
| **Performance** | Poorly defined processes or workflows; inadequate staffing or skill shortages; supplier delays or failures |
| **External** | Regulations; market changes; natural disasters |
| **Organizational** | Scope creep; poor communication; internal organizational changes or priority shifts |

Note the three that catch people out: **supplier delays are Performance**, not External; **scope creep and poor communication are Organizational**, not Performance; **resource misallocation is Financial**, not Organizational.

A project is also **more than the sum of its identified risks**. An **overarching risk** threatens the entire project's success even while every listed risk is being managed. The examples: **lack of stakeholder alignment** (California High-Speed Rail, 2008–), **cultural resistance to change** (Ford's Agile transformation, 2017–2022), and **over-dependence on a vendor or technology** (Australian COVIDSafe app, 2020).

> **Exam focus.** Expect a one-paragraph case and a five-way choice. Ask "where did this originate?" — the money, the technology itself, the way work is performed, the world outside, or the organisation's own structure and behaviour.

`[slides 7-8]`

?check id=rcq-016
?check id=rcq-020

## Risk preference and the risk event graph

**Risk preference** is drawn as a utility-versus-payoff curve. **Risk-averse** is concave (flattening — each extra unit of payoff is worth less, so certainty is preferred); **risk-neutral** is a straight line; **risk-seeking** is convex (steepening).

The **risk event graph** plots two crossing curves over the life cycle. The **chance of a risk occurring is highest at defining and planning** and falls towards delivery; the **cost to fix a risk event is lowest early** and rises steeply towards delivery. That crossover is the whole argument for front-loading risk work.

`[slides 9-10]`

## Step 1 — identify, and how to write a good risk

The **risk management process** has five steps: **1 Risk Identification → 2 Risk Assessment → 3 Risk Prioritization → 4 Risk Response Development → 5 Risk Response Control.**

Identification tools: **brainstorming**, **expert interview (Delphi)**, **historical lessons**, **Risk Breakdown Structure (RBS)** and **SWOT**. The **Delphi method** runs problem identification → selection of experts → open-ended questionnaire → analysis → rating questionnaire → analysis → *is consensus reached?* — if no, loop; if yes, final report. The **RBS** is a hierarchy of risk categories (for an IT project: Business, Technical, Organizational, Project Management, each broken down further). **SWOT** splits internal/external against positive/negative: strengths and weaknesses are internal, opportunities and threats external.

A **good risk statement** explains why, what might happen, and why it matters — **cause → risk event → impact**: *"Because the conference venue has limited backup power (cause), a power outage may occur (event), causing schedule disruption and participant dissatisfaction (impact)."*

Risks are documented in the **risk register**, which "converts uncertainty into managed action" — a table with columns such as No., Rank, Risk Description, Category, Root Cause, Triggers, Potential Responses, Risk Owner, Probability, Impact, Status. Two of those need naming: the **owner** answers *who watches the risk?* and the **trigger** answers *what tells us the risk may be occurring?* (Risk: rain affects outdoor event. Owner: Event Manager. Trigger: forecast shows >70% rain probability.)

`[slides 11-16, 23-24]`

## Steps 2–3 — assess and prioritise

Not all risks deserve equal attention. Assessment uses scenario analysis of **event probability and impact**, a **risk assessment form** (risk event, likelihood, impact, detection difficulty, when), **FMEA**, probability analysis (decision trees, NPV, PERT) and semiquantitative scenario analysis.

**FMEA: Impact × Probability × Detection = Risk Score.** On the worked example — Interface problems 4×4×4 = 64, System freezing 2×5×5 = 50, User backlash 4×3×3 = 36, Hardware malfunctioning 1×5×5 = 25 — note that the *most likely* risk is not automatically the top-scoring one, because detection difficulty is a third multiplier.

Results are plotted on a **probability–impact (risk severity) matrix**: **red zone = major risk, yellow = moderate, green = minor**.

`[slides 17-19]`

## Step 4 — respond, and Step 5 — control

Responses come in matched negative/positive pairs, grouped by what they act on:

| Acts on | Negative (threat) | Positive (opportunity) |
|---|---|---|
| **Cause** | **Avoid** — eliminate the cause (remove activities or people) | **Exploit** — enable the cause (add activities or people) |
| **Outsource** | **Transfer** — find a 3rd party to be accountable or do the work (insurance, fixed-price contract) | **Share** — find a 3rd party to support and share in gain (investors, partnering) |
| **Probability & impact** | **Mitigate** — reduce probability and/or impact (training, prototype) | **Enhance** — increase probability and/or impact (incentives) |
| **Either** | **Accept** — no actionable response | **Accept** |

One risk, four responses — *the keynote speaker may cancel*: **Avoid** = choose another already-confirmed speaker; **Mitigate** = confirm repeatedly and prepare backup content; **Transfer** = contract a speaker agency with a replacement provision; **Accept** = prepare an alternative session if cancellation occurs. Mapped onto the severity matrix: the extreme top-right corner → **avoid**; the high band → **reduce probability** / **reduce impact**; the low-likelihood, low-impact corner → **accept**; the high-impact but unlikely corner → **transfer**.

Money is set aside in two different pots: **contingency reserve** for **identified risks** ("known unknowns") and **management reserve** for **unforeseen uncertainty** ("unknown unknowns").

**Risk control** (step 5) executes the response strategy, monitors triggering events, initiates contingency plans and watches for new risks, supported by a **change management system**: monitoring/tracking/reporting risk, fostering an open environment, repeating identification and assessment, and assigning documented responsibility.

`[slides 20-22, 25-26]`

## Communication planning

Different stakeholders need different information: the **sponsor** wants progress, major risks and budget; the **team** wants tasks, deadlines and issues; the **customer** wants milestones, changes and acceptance; the **supplier** wants specifications and delivery timing. A **power × interest** grid sorts them: high power/high interest = key players, *manage closely*; high power/low interest = *keep satisfied*; low power/high interest = *keep informed*; low power/low interest = *monitor*.

Effective communication = **right information + right person + right time + right method**. The **communication process** runs sender → message → channel → receiver → **feedback**, with **noise/misunderstanding** able to distort any stage.

Planning answers five questions — **WHO** needs information, **WHAT** do they need, **WHEN**/how often, **HOW** (which channel), and **WHO OWNS IT** (who communicates). The plan's components are **Audience · Responsible · Schedule · Method**, written up as a table:

| Stakeholder | Information | Frequency | Method | Owner |
|---|---|---|---|---|
| Sponsor | Status & major risks | Weekly | Meeting | PM |
| Team | Tasks & issues | Daily | Teams | PM |
| Supplier | Delivery status | As needed | Email | Procurement |
| Customer | Milestones | Monthly | Presentation | PM |

`[slides 27-34]`

## Channels, escalation and status reporting

**Synchronous** (real time) = meeting, call, video conference, live chat. **Asynchronous** (different time) = email, report, dashboard, shared document. Rule: **urgent or complex → more synchronous; routine or documented → more asynchronous.**

**Formal** (structured, official) = approvals, status reports, steering committee, contracts — used **for accountability and records**. **Informal** (flexible, casual) = chat, quick call, informal discussion — used **for speed and relationship building**.

**Escalate** when: an **issue exceeds your authority**, a **major risk occurs**, the **baseline is threatened**, or a **stakeholder conflict cannot be resolved**. The ladder runs task owner → project manager → sponsor (provides resources, removes roadblocks) → steering committee (overall direction, key decisions).

A useful status report **directs attention, not just information**: **STATUS** (what happened?) · **VARIANCE** (where are we off plan?) · **RISKS / ISSUES** (what needs attention?) · **DECISION** (what needs management action?).

`[slides 35-38]`

## Quality: definition, planning and metrics

**Quality** is "the degree to which a set of inherent characteristics **fulfils requirements**" (ISO 9000:2000) — **NOT "the best possible"**. The job is to understand the quality levels expected and ensure they are met, which includes **measuring**.

Quality must be **planned before it is inspected**. **Quality planning** defines how good the deliverable must be and how we will know it is acceptable, by answering four questions: **What standard? What metric? What target? How will we check?**

The chain is **Requirement → Quality Standard → Metric → Acceptance Criterion → Check Method**:

| Stage | Meaning | Example |
|---|---|---|
| Requirement | What needs to be achieved | Presentations should start on schedule |
| Quality standard | The level of quality expected | Sessions run on time to respect attendees' time |
| Metric | How performance will be measured | % of sessions starting within 5 minutes of schedule |
| Acceptance criterion | The specific target that must be met | **At least 95%** of sessions start within 5 minutes |
| Check method | How compliance will be verified | Session start-time report from the event system |

A **quality plan** therefore has five columns — Deliverable · Standard · Metric · Target · Check Method — e.g. *Registration | fast service | wait time | <5 min | observation*; *AV system | reliable | major failure | 0 | pre-test*; *Catering | sufficient | meals available | 100% | count*; *Event | satisfaction | rating | ≥4/5 | survey*.

> **Exam focus.** The KPI question shows a proposed quality statement and asks which is correct. A well-formed metric names **what is measured**, a **numeric target** and **how it will be checked** ("survey score ≥ 4/5, measured by post-event survey"). A badly formed one is an adjective with no number and no check: "the system should be fast", "high customer satisfaction", "good-quality catering". Also remember that exceeding the requirement is not higher quality — quality is conformance to requirements.

`[slides 39-44]`

?check id=rcq-044
?check id=rcq-045

## Prevention, QA vs QC, and cost of quality

**Prevent early rather than correct late.** **Prevention** = do it right *before* failure occurs: clear requirements, training, standards, supplier qualification. **Inspection** = check whether the *output* is correct: testing, review, inspection, acceptance testing.

**Quality Assurance** asks *"Are we following the right process?"* — **preventive** and **process-focused**. **Quality Control** asks *"Does the deliverable meet requirements?"* — **detective** and **output-focused**.

**Cost of Quality (CoQ) = Cost of Conformance + Cost of Nonconformance.**

| Cost of conformance (spent to avoid failures) | Cost of nonconformance (spent after failures happen) |
|---|---|
| Prevention cost — costs to avoid defects | Internal failure cost — defects found **before** reaching the customer |
| Appraisal cost — costs to measure and monitor quality | External failure cost — defects found **after** the product reaches the customer |
| Measurement and test equipment costs (part of appraisal) | |

Plotted against quality level, failure cost falls while appraisal and prevention cost rises; total quality cost is a U-shape whose minimum is the **optimal quality level**. Spending more on conformance is worth it only up to that point.

Finally, the three plans interlock: one event may need a **risk** response, a **communication** response and a **quality** response at once — *AV equipment may fail* → escalate immediately to the technical team → restore service within 5 minutes. A good plan **anticipates uncertainty, aligns people and defines acceptable outcomes**.

`[slides 45-52]`
