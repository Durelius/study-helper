# Chulalongkorn University student ID numbers

Working notes on the structure of the 10-digit รหัสประจำตัวนิสิต (student ID), for
building a form validator that will also see IDs from BAScii (Chulalongkorn School
of Integrated Innovation) students across several intake years.

## Segmentation

```
6 9 0 1 0 0 9 2 5 6
└─┬─┘ │ └─┬─┘ │ └─┬─┘
  1    2   3   4   5
```

| # | Positions | Digits | Field | Meaning |
|---|---|---|---|---|
| 1 | 1–2 | 2 | Entry year | Last two digits of the Buddhist-Era (BE) academic year of first enrollment |
| 2 | 3 | 1 | Group/track code | Distinguishes admission group or programme track within that year (exact code meanings not confirmed for Chula — see below) |
| 3 | 4–7 | 4 | Serial number | Sequential number assigned within year + group |
| 4 | 8 | 1 | Check digit | Weighted-sum mod-11 checksum over the other 9 digits |
| 5 | 9–10 | 2 | Faculty/school code | The faculty, school, or institute the student belongs to |

This is a **2 + 1 + 4 + 1 + 2** split, not the 2 + 5 + 1 + 2 split suggested by
eyeballing the numbers. The apparent "5-digit running counter" in the naive
segmentation is actually the 1-digit group code immediately followed by the
4-digit serial number — and separating them is exactly what explains the fourth
ID's outlier value (see below).

- **Confirmed by an official source:** field 5 (faculty code) is in positions
  9–10, and code `56` = สถาบันนวัตกรรมบูรณาการ / **School of Integrated
  Innovation (CSII)** — BAScii's parent school. Confirmed directly from
  Chulalongkorn's own registrar lookup system, `cas.reg.chula.ac.th` (see
  Sources). This is about as official as it gets short of a written policy
  document.
- **Confirmed by verification against all four sample IDs:** field 4 (check
  digit) sits at position 8, and the algorithm below reproduces it exactly for
  all four real IDs given.
- **Corroborated by several unofficial sources + internal arithmetic
  consistency:** field 1 (entry year) is the BE year of first enrollment.
- **Inferred from the four sample IDs only, low confidence:** the exact split
  of field 2 (group code) vs. field 3 (serial number), and what specific group
  code values mean.

## Worked breakdown of the four sample IDs

| ID | Year (1–2) | Group (3) | Serial (4–7) | Check (8) | Faculty (9–10) |
|---|---|---|---|---|---|
| 6901009256 | 69 | 0 | 1009 | 2 | 56 |
| 6901005756 | 69 | 0 | 1005 | 7 | 56 |
| 6901006356 | 69 | 0 | 1006 | 3 | 56 |
| 6758008956 | 67 | 5 | 8008 | 9 | 56 |

- Year `69` = BE 2569 = CE 2026; year `67` = BE 2567 = CE 2024. Consistent with
  the fourth student starting two years earlier than the other three.
- All four end in `56` = School of Integrated Innovation, i.e. all four are
  BAScii students, as expected.
- The three classmates share group code `0` and have close serials (1005,
  1006, 1009) — consistent with being assigned sequential numbers in the same
  admission batch.
- The fourth student has group code `5`, not `0`, and a serial (8008) in a
  completely different numeric range. **This, not a huge intake, is the best
  explanation for the "58008 vs. 01005–01009" anomaly noted in the brief**: it
  isn't one running counter with wildly different values, it's two different
  counters (group `0` vs. group `5`) that happen to be concatenated with the
  year in the naive reading. I could not confirm what group `0` vs. group `5`
  specifically denote (e.g., different admission round, quota vs. international
  track, direct-entry vs. transfer) — flagging as unconfirmed rather than
  guessing.

## Check digit algorithm — found and verified

**Source:** an unofficial, unattributed GitLab snippet ("Chula Student ID
Check Digit"), the only place a check-digit formula for Chula IDs could be
found in this research. No official Chula document describing a check digit
was located. I verified the algorithm by hand against all four real sample
IDs — it matches all four, which is strong practical confirmation even though
the written source itself is unofficial.

**Algorithm** (for the 10-digit format used from entry year 62 onward):

```
weights = [3, 1, 7, 2, 1, 3, 7, 0, 7, 3]   # applied to digits 1..10
checksum = ( Σ digit[i] * weight[i] for weight[i] != 0 )  mod 11  mod 10
digit[8] (the check digit) must equal checksum
```

Digit 8 itself is excluded from its own sum (weight 0); digits 9–10 (the
faculty code) *are* included in the sum with weights 7 and 3.

**Worked example — 6901009256:**

digits: 6 9 0 1 0 0 9 2 5 6 (positions 1–10)
weights: 3 1 7 2 1 3 7 0 7 3

```
6×3 + 9×1 + 0×7 + 1×2 + 0×1 + 0×3 + 9×7 + 5×7 + 6×3
= 18 + 9 + 0 + 2 + 0 + 0 + 63 + 35 + 18
= 145
145 mod 11 = 2      (11 × 13 = 143)
2 mod 10   = 2
→ checksum = 2, digit[8] = 2 ✓ match
```

**Verification against all four samples:**

| ID | Sum | mod 11 | mod 10 | Actual check digit | Match? |
|---|---|---|---|---|---|
| 6901009256 | 145 | 2 | 2 | 2 | yes |
| 6901005756 | 117 | 7 | 7 | 7 | yes |
| 6901006356 | 124 | 3 | 3 | 3 | yes |
| 6758008956 | 185 | 9 | 9 | 9 | yes |

All four verify exactly. This is a small sample (n=4), but a 4-for-4 match on
an independently-sourced formula is meaningful — a wrong formula would need a
1-in-10,000 coincidence to pass all four by chance.

The same snippet also documents an **older 8-digit format** for entry years
before 62 (BE 2562 / CE 2019): 2 (year) + 5 (group+serial) + 1 (check digit),
weights `[3,1,7,2,1,3,7]`, no embedded faculty code, check digit as the last
(8th) digit. It claims the switch to the current 10-digit format happened at
entry year 62. This is a useful coincidence to flag, not a proof: BAScii
itself only launched in 2019 (BE 2562, i.e. "62"), so **every BAScii student
ID should be in the newer 10-digit format** — but this format-change claim
itself is corroborated only by the same single unofficial source, not an
official Chula announcement.

## Entry year corroboration

Beyond the raw samples, the entry-year reading is supported by Chulalongkorn's
own "รุ่น" (cohort/generation) numbering, found in official social posts:
generation number = (last two digits of the BE entry year) + 40. Example
found in search results: an official Chulalongkorn Student Affairs post
congratulating new cohort **"Chula 109"** — and 69 + 40 = 109, matching entry
year 69 (BE 2569 / CE 2026) exactly. This is arithmetically self-consistent
with Chulalongkorn's founding year of BE 2460 (generation 109 → BE
2460 + 109 = 2569). This is corroboration, not an official single-page
statement of the ID format itself.

## What a validator can safely enforce

| Check | Enforce? | Why |
|---|---|---|
| Exactly 10 digits, numeric only | Yes | Confirmed format for entry year ≥ 62; BAScii has never had an intake before 2562/"62", so all BAScii IDs will be 10 digits |
| Digits 1–2 form a plausible entry year | Yes, loosely | Use a generous range (e.g. current BE year down to ~55–60) rather than a tight window, since the app may see several intake cohorts and formats may shift again |
| Check digit (position 8) matches the mod-11 algorithm above | Yes, but as a soft check | Verified against all 4 real samples, but the only source is an unofficial snippet — treat a mismatch as a warning ("this doesn't look right, please double-check") rather than a hard rejection, in case of format drift or edge cases not covered here |
| Digits 9–10 are a known faculty code (e.g. `56` for BAScii-only tooling) | Optional | The code list from `cas.reg.chula.ac.th` is official but the retrieved listing had gaps (only ~29 codes were present: 01, 02, 20–40, 51, 53, 55, 56, 58, 99) — treat it as a lookup/allow-list for known codes, not a complete enumeration to hard-reject unknown ones, unless scoping the tool to BAScii only (require `56`) |
| Digit 3 (group code) is a specific value | **No** | Only two values observed across 4 samples (0 and 5); meaning unconfirmed. Accept any digit 0–9. |
| Digits 4–7 (serial) fall in a specific range | **No** | Only 4 data points spanning two very different ranges (1005–1009 and 8008); no basis to bound this. Accept any 4-digit value. |

**What would change these conclusions:** a counter-example BAScii ID with an
8-digit format, a check digit that fails the mod-11 test, or a faculty code
other than `56` would all be reason to revisit the corresponding claim above.
The group-code and serial-number claims are the weakest in this document and
should be treated as provisional pending more samples.

## Sources

- **Official — Chulalongkorn Office of the Registrar, faculty code lookup**:
  `https://cas.reg.chula.ac.th/cu/general/PersonalInformation/Faculty/IndexDisplayFaculty.html`
  (frames to `https://cas.reg.chula.ac.th/servlet/com.dtm.chula.general.servlet.FacultyServlet`).
  Confirms code `56` = สถาบันนวัตกรรมบูรณาการ / School of Integrated
  Innovation, and that faculty codes occupy the last two digits of the ID
  (per accompanying page text: "หลักที่ 9-10 ของรหัสประจำตัวนิสิต คือรหัสคณะ").
- **Official — Chulalongkorn / CSII program pages** confirming BAScii sits
  under the School of Integrated Innovation (CSII):
  `https://www.chula.ac.th/program/arts-and-science-in-integrated-innovation/`,
  `https://www.chula.ac.th/en/departments/school-of-integrated-innovation/`,
  `https://csii.chula.ac.th/bascii/`.
- **Unofficial, but verified by hand against all 4 sample IDs** — check-digit
  algorithm and format-change-at-year-62 claim: GitLab snippet "Chula Student
  ID Check Digit", `https://gitlab.com/-/snippets/1982147`.
- **Corroborating (unofficial, official-adjacent)** — Chulalongkorn Student
  Affairs Facebook post referencing cohort "Chula 109" (consistent with entry
  year 69 via the +40 generation formula):
  `https://www.facebook.com/StudentAffair.CU/` (post found via search, title
  "ขอแสดงความยินดีกับนิสิตใหม่ #Chula109").
- **Explicitly ruled out / do not reuse**: a page titled "รหัสประจำตัวนิสิต 10
  หลัก ที่ควรทราบ" at `reg.mcu.ac.th` looked like a strong hit in search but
  is from **Mahachulalongkornrajavidyalaya University (MCU)**, a separate
  Buddhist university with a similar-sounding name — not Chulalongkorn
  University. Its content (and any digit-3 "0/1/3/4/7/8 = undergrad/grad ×
  science/social-science" breakdown attributed to it) was excluded from this
  document as misattributed, not as evidence about Chula's own ID format.
- Several other Chula pages were checked for a written description of the ID
  format (`reg.chula.ac.th` student handbook and card pages, the `sa.chula.ac.th`
  short manual, the Faculty of Education จท99 handbook, `web.reg.chula.ac.th`
  Q&A page) but none contained an explicit breakdown of the digit structure —
  only the registrar's faculty-code lookup tool and the unofficial GitLab
  snippet yielded concrete, checkable content.
