---
id: project-monitoring
title: "Project Monitoring, Control & Earned Value"
deck: "new 10"
slides: 39
summary: "How to tell whether a project is still following its approved baseline: scope, schedule and cost tracking, change control, and earned value as the test of performance."
---

## Execution, monitoring and control are three different jobs

Execution is doing the work: developing the system, coordinating people. Monitoring is finding out what is happening.
Control is deciding what to do about it. The deck's example: execution builds the inventory system, monitoring notices that
testing is two weeks behind, and control reallocates testing resources. Monitoring detects deviations, and control decides
the response and carries it out. `[slides 4-5]`

The control cycle runs the same way every time: measure against the baseline, compare, decide, act, and re-baseline if
the change is approved. `[slide 6]`

## The baseline is the reference point

A **baseline** is an approved reference used to measure and control performance. There are three: the **scope baseline**
(the approved deliverables and work), the **schedule baseline** (activities and dates), and the **cost baseline** (the
time-phased budget). Monitoring always asks how actual performance compares with these. `[slide 7]`

> **Exam focus.** Without an approved baseline there is nothing to measure against. A forecast is not a baseline, and
> neither is a plan that nobody has approved.

## Scope: completed is not accepted

Scope monitoring asks three questions: what was originally approved, what has been completed and accepted, and has anyone
asked for additional or different work. `[slide 8]`

**Scope progress** measures delivery against the existing plan. **Scope change** modifies the approved work or
requirements. The two are separate measurements. `[slide 9]`

The FreshMart store rollout shows why. Of 25 planned stores, 15 are deployed and only 13 are accepted, so accepted
progress is 13, not 15. `[slides 11-12]`

Scope creep is additional work without authorisation. Controlled change is documented, its impact is assessed, and the
baseline is updated only if it is approved. Scope creep makes the baseline unreliable, and controlled change does not.
`[slides 13-14]`

A **scope change log** traces each request from proposal to decision, with its scope impact. In the deck's log, one
request is pending, one is rejected and one is approved and implemented. Only the approved change becomes part of the
scope. `[slide 15]`

## Schedule: variance is measured in time

Schedule monitoring compares actual or forecast progress with the schedule baseline. It asks which activities have started,
finished or been delayed, whether the delays will affect milestones, and whether the team can recover without moving the
deadline. `[slide 16]`

For completed activities use the actual finish; for unfinished ones use the forecast. A schedule variance in time is the
forecast finish minus the baseline finish. Zero means on schedule, positive means late and negative means early. The
deck's example is a 2-week slip (baseline week 9, forecast week 11). `[slides 17-18]`

Milestones have zero duration and mark a significant event, such as "Design approved" or "Project handover." A milestone
that moves is the early warning sign. `[slide 19]`

> **Exam focus.** A delay only moves the project if it hits the critical path. Critical activities have zero total float,
> so their delay moves the finish unless recovered. Non-critical activities absorb delays within their float. The critical
> path can change during execution, so check it again after every significant delay. `[slides 20-21]`

The deck's example: the critical path is A→B→E→G. If B is delayed by 2 days and nothing is recovered, the project duration
goes from 20 to 22 days. `[slide 21]`

When a schedule slips, managers choose a response, and each has a trade-off. They can **reallocate resources** (other work
loses capacity), **fast-track** by overlapping activities (higher rework risk), **crash** by adding resources to critical
activities (higher cost), **reduce nonessential effort** (limited recovery), or **submit a change request** for a revised
date (requires approval). `[slide 22]`

Keep three dates apart: the **baseline finish** (the original target), the **forecast finish** (the current prediction) and
the **revised baseline finish** (the new target, which only exists once the change is approved). In the deck's example the
forecast is week 18, but the revision is still pending. `[slide 23]`

## Cost: spending is not performance

Cost monitoring asks whether spending and forecast costs stay within the approved budget. Four figures carry the analysis:
**BAC** (Budget at Completion, the authorised total), **planned expenditure** (spending scheduled up to the reporting date),
**AC** (Actual Cost) and **EAC** (Estimate at Completion, the forecast total). The question is whether the project is
spending the right amount for the work completed. `[slide 24]`

The FreshMart month-4 table shows planned and actual spending by category. Total spending matches exactly at ฿2.8M, yet
the deck asks whether the project is financially on track. The answer is not yet, because spending on plan says nothing
about whether the planned work was done. `[slide 25]`

The sharpest example compares two ฿1 million projects that both planned 50% of work by today. Project A has verified 50%
and spent ฿450,000. Project B has verified only 30% and spent ฿350,000. Project B spent less but delivered far less, so it
is the weaker performer. `[slide 26]`

## Earned value integrates scope, schedule and cost

**Earned Value Management (EVM)** combines the three into one measure. It needs a budgeted baseline and a credible way of
measuring completed work. `[slide 27]`

The terms are **PV** (value of planned work), **EV** (value of work completed), and **AC** (cost actually incurred). In the
FreshMart month-4 example, PV is ฿2.8M, EV is ฿2.4M and AC is ฿2.8M. The project should have completed ฿2.8M of work, has
completed ฿2.4M, and has spent ฿2.8M. It is both behind schedule and over budget. `[slide 28]`

**Schedule Variance** is SV = EV − PV, measured in budget-value units, not calendar days. **Cost Variance** is CV = EV − AC.
FreshMart gives SV = 2.4 − 2.8 = −฿0.4M and CV = 2.4 − 2.8 = −฿0.4M, both unfavourable. `[slide 29]`

**SPI** = EV / PV and **CPI** = EV / AC. Both are 1.00 on plan. FreshMart gives 0.86 on each: behind schedule in
earned-value terms, and cost inefficient. `[slide 30]`

If current cost efficiency continues, the final cost forecast is **EAC = BAC / CPI**. FreshMart has BAC ฿6.0M and CPI
0.857, so EAC = 6.0 / 0.857 = ฿7.0M, a potential overrun of ฿1M. The deck asks whether management should approve the extra
฿1M automatically. It should not: the forecast assumes today's efficiency holds, so the overrun needs a change decision.
`[slide 31]`

Scope, schedule and cost are linked, so a change in one affects the others. A stakeholder request for five extra stores
changes the scope, and it also moves the schedule and the cost. `[slide 32]`

## Deviation or change: the response differs

A task that is two weeks late is a schedule deviation, and the response is to investigate or recover. Cost above the value
of work is a cost-performance problem, and the response is to analyse its causes. A customer request for five more stores
is a proposed scope change, so it is evaluated as a change request. A sponsor's approved new deadline is an approved
baseline change, and the schedule baseline is updated. `[slide 33]`

A **deviation** is actual or forecast performance that differs from the baseline. A **change request** is a proposal to
modify an approved project. Only an approved change moves the baseline. `[slide 33]`

## Integrated change control

Any change follows the same seven steps: identify the deviation or proposed change, record it as an issue or change
request, assess scope, schedule, cost, quality and risk impacts, develop response alternatives, submit for approval when
required, update the affected baselines once approved, and then communicate, implement and monitor. `[slide 34]`

Governance settles who decides. The project team updates progress and raises issues. The project manager monitors
performance and starts corrective action. The sponsor resolves escalated issues and approves authorised changes. A
steering committee or change control board approves major changes within its delegated authority. `[slide 35]`

A status report should show performance, exceptions, forecasts and the decisions required. A report that only says what
is wrong is incomplete. It should say what decision management needs to make. `[slide 36]`

For the group project, the progress update has two parts. The updated Gantt chart shows the approved baseline against actual
or forecast progress, the completion percentages, the delayed activities and milestones, and the reasons for significant
differences. The status and action plan covers scope (what is complete, what changes were requested), schedule (on time,
which activities are delayed), cost (planned against actual, where applicable), issues, actions, and the next milestone.
`[slides 37-38]`

## Wrap-up

Approved baselines are the reference for monitoring. Scope tracking separates completed, accepted and changed work.
Schedule tracking compares actual and forecast dates with baseline dates. Cost tracking compares spending with the value
of completed work. Scope, schedule and cost changes must be assessed together, and governance makes sure the right
authority approves them. `[slide 39]`

?check id=pmo-001
?check id=pmo-003
?check id=pmo-009
?check id=pmo-011
