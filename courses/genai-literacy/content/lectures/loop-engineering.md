---
id: loop-engineering
title: "Loop Engineering and Agents"
deck: "Nov 13 quiz"
slides: 0
summary: "Designing the repeated cycle an agent runs, and the part the human keeps: the goal, the criteria and the stop."
---

## The loop, not the single prompt

A single prompt gets one answer. An **agent loop** repeats a cycle: **plan**, **act** (call a tool or write something), **observe** the result, and **check** it against the goal. The cycle continues until a **stop condition** says it is done. Loop engineering is the design of that cycle: what the agent may do, how it judges its own output, and when it must stop.

> **Exam focus.** A loop is a repeated cycle with a stop rule. One prompt and one answer is not a loop.

`[Nov 13 quiz, loop engineering]`

## What the human keeps

The quiz's answer to "What is the human's role?" is that the human **defines the goal and workflow**. The same logic gives three responsibilities that belong to the human in loop engineering:

- **Define the goal and the workflow.** The shape of the job and what success means.
- **Set the evaluation criteria.** The test for when output counts as good. An agent that judges its own work with no outside check can declare victory too early.
- **Decide the stop condition.** The rule that ends the loop.

The one thing the human does *not* do is dictate every step. Dictating each step removes the loop and the agent's ability to adapt.

> **Exam focus.** Human sets goal, criteria and stop. Agent chooses the steps inside them.

## Stop conditions

A good stop condition combines a **success test** with a **hard limit**. "Stop when all tests pass or after a set number of rounds, whichever comes first" is the safe form. "Run until the agent says it is finished" trusts the agent's own claim, which is the failure the criteria exist to prevent. A limit on iterations, sometimes called an iteration cap, stops an endless loop from running up cost.

## Building a loop that works

- **Observe before you act again.** A loop that never looks at the result repeats its mistake. In the common case of an agent that keeps undoing its own edits, the fix is a check that compares the result with the goal, plus a stop rule.
- **Log each step.** A human must be able to review what the agent did.
- **Set limits up front.** Iteration caps and cost caps belong in the design, not added after a runaway run.

## Loop versus prompt

Prompt engineering shapes one request. Loop engineering designs the cycle around one or more model calls, including the checks that decide whether to continue. Both matter; a loop is only as good as the prompts inside it.

?check id=le-001
?check id=le-002
?check id=le-004
?check id=le-007
