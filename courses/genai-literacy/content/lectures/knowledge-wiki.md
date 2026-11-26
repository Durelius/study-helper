---
id: knowledge-wiki
title: "RAG and the LLM Wiki"
deck: "Nov 13 quiz"
slides: 0
summary: "How retrieval gives a model your own knowledge, and the three operations that keep an LLM Wiki working."
---

## Retrieval-augmented generation

A model knows only what it was trained on. **Retrieval-augmented generation (RAG)** lets it use your own material without retraining: documents are split into **chunks**, each chunk is turned into an **embedding**, and the embeddings go into a **vector store**. When a question arrives, the system embeds the question, retrieves the chunks closest to it, places them in the prompt, and only then generates.

The steps, in order:

1. Chunk and embed the documents, and store them.
2. Embed the question.
3. Retrieve the closest chunks from the store.
4. Add them to the prompt.
5. Generate the answer from that context.

> **Exam focus.** RAG lets you include local knowledge in your AI. It uses embeddings to find relevant passages, and the model generates from what it retrieves.

`[Nov 13 quiz, RAG]`

## What RAG is not

RAG does not change the model's weights, so it does not memorise your documents. It does not browse the open web unless a search step is added. It does not guarantee that the knowledge base is free of contradictions, because the sources themselves can disagree. RAG reduces some errors and makes answers traceable to sources, but it does not remove them.

## Embeddings

An **embedding** is a vector that represents a text's meaning, arranged so that similar meanings sit close together. Retrieval works by finding the nearest vectors to the question's vector. Without embeddings there is no similarity search, which is why the quiz names them as part of RAG.

## RAG versus fine-tuning

Use RAG when the knowledge changes, must be cited, or is too large for a model to absorb. Updating a document store is cheap; retraining the model is not. Fine-tuning changes the weights and suits learning a style or a narrow skill.

## The LLM Wiki

An **LLM Wiki** is a set of interlinked pages that an LLM maintains from your sources. It runs three recurring operations:

- **Ingest.** A new source is read and folded into the wiki, updating the pages it bears on. This is the write path.
- **Query.** A question is answered from the existing pages, with the answer grounded in them.
- **Lint.** The wiki is checked for contradictions, gaps and stale claims. This keeps it healthy as sources accumulate.

The operations are distinct. Ingest and query are the two everyday uses; lint is the maintenance step that stops a growing wiki drifting apart. Rendering and coding are not wiki operations.

> **Exam focus.** Ingest, Query, Lint. Know which one writes, which one answers, and which one audits.

`[Nov 13 quiz, LLM Wiki]`

?check id=kw-001
?check id=kw-002
?check id=kw-006
?check id=kw-008
