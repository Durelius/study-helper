---
id: context-engineering
title: "Context Engineering"
deck: "Nov 13 quiz"
slides: 0
summary: "What is inside a model's context, why more is not better, and how to keep the window useful."
---

## What the context is

The **context** is everything the model is given for one answer: the system prompt, the current user prompt, the conversation history, the tool descriptions, and any files the model has been handed in this session. The model's training weights are not context, and neither is anything it has never received. A file sitting on your hard drive that you never pasted or retrieved cannot influence the reply.

> **Exam focus.** Anything not given to the model is not in its context, whatever the user knows.

`[Nov 13 quiz, Q: Which information is not part of context]`

## Why context engineering exists

**Context engineering** is the discipline of choosing what goes into the window, what stays, what is summarised, and what is fetched only when it is needed. Its purpose is to keep the context useful and to reduce the problems that irrelevant or excessive context causes. It does not guarantee correct answers, and it is not the same as making an agent run unattended.

`[Nov 13 quiz, Q: Why context engineering]`

## Bigger is not better

The **context window** is the amount of text (measured in tokens) that a model can hold at once. A larger window lets more fit, but it does not make the model use every token equally well. Information in the middle of a long context is often attended to less, and irrelevant material distracts.

Two failure modes come up again and again:

- **Context rot**: answer quality falls as the window fills with stale or off-topic material.
- **Distraction**: the model drifts toward facts that were in the clutter but are not what the question asked about.

Symptoms include slower and costlier replies, ignored instructions buried deep in the window, and answers that wander.

## Practical moves

1. **Retrieve, don't paste.** For a 300-page report and a question about one section, include the relevant section only.
2. **Summarise old turns.** Keep the task-relevant facts from a long conversation and drop the rest.
3. **Keep tool descriptions tight.** The model reads them every turn to choose a tool, so vague or duplicate descriptions cost attention.
4. **Put the instruction where it will be read.** Important rules belong near the task, not buried in a long history.

> **Exam focus.** The right answer to a long-context problem is almost always to **curate**, not to enlarge the window.

`[Nov 13 quiz, context engineering]`

?check id=ce-001
?check id=ce-002
?check id=ce-003
?check id=ce-009
