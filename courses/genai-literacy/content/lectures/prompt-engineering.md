---
id: prompt-engineering
title: "Prompt Engineering: Roles, Shots and Structure"
deck: "Nov 13 quiz"
slides: 0
summary: "What a prompt is made of, and what the role, the examples and the output format each do."
---

## A prompt is a set of instructions

A **prompt** is the text you give a model on one turn. A well-structured prompt contains three things: a **role** (who the model should be), a **task** (what it should do) and an **output format** (what shape the answer takes). Examples can go in too. A prompt that has only a task, such as "Estimate the price", leaves the model to guess the rest.

> **Exam focus.** Role sets the expertise and behaviour. It does not change the context window, the file format or saving.

`[Nov 13 quiz, Q: Role]`

## Role

Writing "You are a Bangkok property analyst" tells the model which perspective, vocabulary and level of detail to answer from. It is a steer, not a guarantee of accuracy. The role and the task are different parts of the prompt: the role says *who*, the task says *what*. Keeping them separate lets you change one without rewriting the other.

## Zero-shot, few-shot

The number of **shots** is the number of worked examples in the prompt.

- **Zero-shot**: the instruction alone, with no examples. The model relies on what it already knows.
- **Few-shot**: a small number of input-output examples placed before the new input. The model copies the pattern, including format and tone.

Few-shot changes only the input. The weights are not updated, and the examples are forgotten when the conversation ends. Changing weights is fine-tuning, a different technique.

> **Exam focus.** Few-shot is "a small number of examples before the task". It is not "very few words", "the same prompt many times", or "no examples".

`[Nov 13 quiz, Q: Few-shot]`

## Chain-of-thought and output format

**Chain-of-thought** asks the model to show its intermediate steps before the answer. For multi-step problems the steps often improve the result.

The **output format** specification, such as "reply with one number in baht", fixes the shape of the answer and makes it easy to check. It is the cheapest structure to add and the one most often missing.

## Putting it together

The strongest prompt in a quiz will have all of the above: a role, a few examples, a clear task and a format. The weakest is a bare request like "Price?", which gives the model nothing to anchor on.

`[Nov 13 quiz, prompt comparison]`

?check id=pe-001
?check id=pe-002
?check id=pe-004
?check id=pe-012
