"""Figures for the Economic Concepts final-quiz handout.

Edit the data lists below and rerun:  python3 figures.py
Each figure is written as figures/<name>.svg and a transparent figures/<name>.png (300 dpi).
"""
import os
import numpy as np
import matplotlib
matplotlib.use("Agg")
import matplotlib.pyplot as plt

# ---------------------------------------------------------------- data ---------
# Oligopoly: the packet's market (P = A - Q, MC = C) and the four outcomes.
MARKET = {"A": 339, "C": 147}
STRUCTURES = [  # name, share of (A - C) produced in total, colour
    ("Bertrand", 1.0, "#16a34a"),
    ("Stackelberg", 0.75, "#2563eb"),
    ("Cournot", 2 / 3, "#7c3aed"),
    ("Collusion", 0.5, "#dc2626"),
]

LABEL_POS = {  # where each point's name sits: offset in points, alignment
    "Collusion": ((8, 6), "left", "bottom"),
    "Cournot": ((-8, 6), "right", "bottom"),
    "Stackelberg": ((9, 3), "left", "bottom"),
    "Bertrand": ((8, 8), "left", "bottom"),
}

# Solow: y = k^ALPHA, saving rates, depreciation.
SOLOW = {"alpha": 0.5, "s1": 0.3, "s2": 0.45, "delta": 0.1, "delta2": 0.15}

# Market for ideas: horizontal fragmentation levels (0 = unified, 1 = fragmented).
IDEAS_LEVELS = [  # label, fragmentation, colour
    ("Africa", 0.86, "#6b7280"),
    ("Europe", 0.68, "#2563eb"),
    ("Europe's audience\n(Republic of Letters)", 0.30, "#2563eb"),
    ("China", 0.12, "#ea580c"),
]

# Surplus: the general picture (straight supply and demand), drawn on a 0-10 grid.
SURPLUS_GENERAL = {"demand_top": 10, "supply_bottom": 2, "q_restricted": 2.6}
# Colours of the three areas: consumer surplus, producer surplus, deadweight loss.
SURPLUS_COLOURS = {"CS": "#2563eb", "PS": "#16a34a", "DWL": "#dc2626"}

# Computers example (slide 88): N computers = N workers, rates per period.
COMPUTERS = {"N": 10, "delta": 0.1, "n": 0.2, "g": 1.0}

# ---------------------------------------------------------------- style --------
INK, MUTED, GRID = "#1f2937", "#6b7280", "#e5e7eb"
BLUE, ORANGE, GREEN, RED, PURPLE = "#2563eb", "#ea580c", "#16a34a", "#dc2626", "#7c3aed"
plt.rcParams.update({
    "font.family": ["Avenir Next", "Arial Unicode MS"], "font.size": 10, "text.color": INK,
    "axes.edgecolor": INK, "axes.labelcolor": INK, "axes.linewidth": 1.0,
    "xtick.color": MUTED, "ytick.color": MUTED, "axes.spines.top": False,
    "axes.spines.right": False, "svg.fonttype": "none",
})
OUT = os.path.join(os.path.dirname(os.path.abspath(__file__)), "figures")
os.makedirs(OUT, exist_ok=True)


def save(fig, name):
    fig.savefig(os.path.join(OUT, name + ".svg"), transparent=True, bbox_inches="tight")
    fig.savefig(os.path.join(OUT, name + ".png"), dpi=300, transparent=True, bbox_inches="tight")
    plt.close(fig)


def bare(ax, xlabel, ylabel):
    ax.set_xticks([]); ax.set_yticks([])
    ax.set_xlabel(xlabel, loc="right"); ax.set_ylabel(ylabel, loc="top")


def f(k):
    return k ** SOLOW["alpha"]


def kstar(s, d):
    return (s / d) ** (1 / (1 - SOLOW["alpha"]))


# ---------------------------------------------------------------- 1 demand -----
def fig_demand():
    A, C = MARKET["A"], MARKET["C"]
    G = A - C
    fig, ax = plt.subplots(figsize=(6.2, 3.9))
    q = np.linspace(0, A, 50)
    ax.plot(q, A - q, color=INK, lw=2)
    ax.axhline(C, color=MUTED, lw=1.4, ls="--")
    ax.text(A * 0.97, C - 8, f"MC = {C}", color=MUTED, ha="right", va="top")
    ax.text(18, A - 18 + 4, f"Demand  P = {A} − Q", va="bottom")
    for name, share, col in STRUCTURES:
        Q = share * G
        P = A - Q
        ax.plot([Q, Q], [0, P], color=col, lw=0.9, ls=":")
        ax.plot([0, Q], [P, P], color=col, lw=0.9, ls=":")
        ax.scatter([Q], [P], color=col, s=42, zorder=5)
        off, ha, va = LABEL_POS[name]
        ax.annotate(name, (Q, P), xytext=off, textcoords="offset points", color=col,
                    fontsize=8.5, fontweight="demibold", ha=ha, va=va)
    rows = "\n".join(f"{n:<12} Q = {s * G:>4g}   P = {A - s * G:>4g}" for n, s, _ in STRUCTURES)
    ax.text(0.98, 0.97, rows, transform=ax.transAxes, ha="right", va="top", fontsize=8.5,
            family="Menlo", color=INK, bbox=dict(boxstyle="round,pad=0.5", fc="white", ec=GRID))
    ax.set_xlim(0, A); ax.set_ylim(0, A + 10)
    ax.set_xlabel("Total output Q (rides / min)"); ax.set_ylabel("Price P (THB)")
    ax.set_xticks([0] + [s * G for _, s, _ in STRUCTURES] + [A])
    ax.set_yticks([0, C, A])
    ax.tick_params(labelsize=8)
    save(fig, "demand")


# ---------------------------------------------------------------- 2 ideas ------
def fig_ideas():
    fig, ax = plt.subplots(figsize=(6.2, 4.4))
    def aud(x):  # downward: the ideas an audience of this size supports
        return 0.02 + 0.17 / x

    def rsk(x):  # upward: the ideas this level of persecution permits
        return 0.05 + 1.25 * x

    fa = np.linspace(0.16, 0.98, 200)
    fr = np.linspace(0.04, 0.84, 200)
    ax.plot(aud(fa), fa, color=INK, lw=2)
    ax.plot(rsk(fr), fr, color=INK, lw=2)
    ax.text(aud(0.16), 0.20, "Size of audience\n(demand)", fontsize=9, ha="right", va="bottom")
    ax.text(rsk(0.84) + 0.01, 0.84, "Risk of\npersecution\n(supply)", fontsize=9, va="center")

    for label, y, col in IDEAS_LEVELS:
        ax.axhline(y, color=col, lw=0.8, ls="--", alpha=0.7)
        ax.text(0.01, y + 0.012, label, color=col, ha="left", va="bottom", fontsize=8, fontweight="demibold")
    pts = [
        ("", aud(0.86), 0.86, MUTED, (6, 4)),
        ("E″", aud(0.68), 0.68, BLUE, (6, 4)),
        ("E₁", rsk(0.68), 0.68, BLUE, (6, -12)),
        ("E′", rsk(0.30), 0.30, BLUE, (-6, 6)),
        ("E₂", aud(0.30), 0.30, BLUE, (6, 4)),
        ("", rsk(0.12), 0.12, ORANGE, (6, -12)),
    ]
    for name, x, y, col, off in pts:
        ax.scatter([x], [y], color=col, s=40, zorder=5)
        ax.annotate(name, (x, y), xytext=off, textcoords="offset points", color=col, fontsize=9, fontweight="demibold")
    ax.annotate("", xy=(rsk(0.68) - 0.01, 0.655), xytext=(aud(0.30) + 0.01, 0.32),
                arrowprops=dict(arrowstyle="-|>", color=RED, lw=1.6))
    ax.text(0.80, 0.42, "Europe gets both:\nlow persecution (E₁)\n+ a unified audience (E₂)", color=RED, fontsize=8.5)
    ax.set_xlim(0, 1.25); ax.set_ylim(0, 1.0)
    bare(ax, "Number of ideas →", "Fragmentation index ↑")
    save(fig, "ideas")


# ---------------------------------------------------------------- 3 solow build -
def fig_solow_build():
    s, d = SOLOW["s1"], SOLOW["delta"]
    ks = kstar(s, d)
    k = np.linspace(0, ks * 1.7, 300)
    fig, ax = plt.subplots(figsize=(6.2, 3.9))
    ax.plot(k, f(k), color=INK, lw=2, label="Output  y = f(k)")
    ax.plot(k, s * f(k), color=BLUE, lw=2, label="Investment  s·f(k)")
    ax.plot(k, d * k, color=ORANGE, lw=2, label="Depreciation  δk")
    ax.axvline(ks, color=MUTED, lw=0.8, ls=":")
    ax.scatter([ks], [s * f(ks)], color=INK, s=40, zorder=5)
    ax.annotate("Steady state:\ns·f(k) = δk, Δk = 0", (ks, s * f(ks)), xytext=(-30, 40), textcoords="offset points", fontsize=8.5, ha="right",
                arrowprops=dict(arrowstyle="-", color=MUTED, lw=0.8))
    kc = ks * 0.45
    ax.annotate("", xy=(kc, f(kc)), xytext=(kc, s * f(kc)), arrowprops=dict(arrowstyle="<->", color=GREEN, lw=1.3))
    ax.text(kc + 0.2, (f(kc) + s * f(kc)) / 2, "c  (consumption)", color=GREEN, fontsize=8.5, va="center")
    ax.annotate("", xy=(kc, s * f(kc)), xytext=(kc, 0), arrowprops=dict(arrowstyle="<->", color=BLUE, lw=1.3))
    ax.text(kc + 0.2, s * f(kc) / 2, "i", color=BLUE, fontsize=9, va="center")
    ax.annotate("k grows →", (ks * 0.55, 0.12), color=GREEN, fontsize=8.5, fontweight="demibold")
    ax.annotate("← k shrinks", (ks * 1.25, 0.12), color=RED, fontsize=8.5, fontweight="demibold")
    ax.set_xticks([ks]); ax.set_xticklabels(["k*"]); ax.set_yticks([])
    ax.set_xlim(0, ks * 1.7); ax.set_ylim(0, f(ks * 1.7) * 1.05)
    ax.set_xlabel("Capital per worker k", loc="right")
    ax.legend(frameon=False, fontsize=8.5, loc="upper left")
    save(fig, "solow_build")


# ---------------------------------------------------------------- 4 saving -----
def fig_solow_saving():
    s1, s2, d = SOLOW["s1"], SOLOW["s2"], SOLOW["delta"]
    k1, k2 = kstar(s1, d), kstar(s2, d)
    k = np.linspace(0, k2 * 1.3, 300)
    fig, ax = plt.subplots(figsize=(6.2, 3.6))
    ax.plot(k, d * k, color=ORANGE, lw=2, label="δk")
    ax.plot(k, s1 * f(k), color=BLUE, lw=2, alpha=0.45, label="s₁·f(k)  (before)")
    ax.plot(k, s2 * f(k), color=BLUE, lw=2, label="s₂·f(k)  (after)")
    for kk, lab in [(k1, "k₁*"), (k2, "k₂*")]:
        ax.axvline(kk, color=MUTED, lw=0.8, ls=":")
        ax.scatter([kk], [d * kk], color=INK, s=36, zorder=5)
    ax.annotate("", xy=(k2, 0.25), xytext=(k1, 0.25), arrowprops=dict(arrowstyle="-|>", color=GREEN, lw=1.6))
    ax.text((k1 + k2) / 2, 0.32, "transition:\nfaster growth", color=GREEN, ha="center", fontsize=8.5)
    ax.set_xticks([k1, k2]); ax.set_xticklabels(["k₁*", "k₂*"]); ax.set_yticks([])
    ax.set_xlabel("Capital per worker k", loc="right")
    ax.legend(frameon=False, fontsize=8.5, loc="upper left")
    save(fig, "solow_saving")


# ---------------------------------------------------------------- 5 delta / n --
def fig_solow_breakeven():
    s, d1, d2 = SOLOW["s1"], SOLOW["delta"], SOLOW["delta2"]
    k1, k2 = kstar(s, d1), kstar(s, d2)
    k = np.linspace(0, k1 * 1.35, 300)
    fig, ax = plt.subplots(figsize=(6.2, 3.6))
    ax.plot(k, s * f(k), color=BLUE, lw=2, label="s·f(k)")
    ax.plot(k, d1 * k, color=ORANGE, lw=2, alpha=0.45, label="(δ + n)k  before")
    ax.plot(k, d2 * k, color=ORANGE, lw=2, label="(δ + n)k  after δ or n rises")
    for kk in (k1, k2):
        ax.axvline(kk, color=MUTED, lw=0.8, ls=":")
        ax.scatter([kk], [s * f(kk)], color=INK, s=36, zorder=5)
    ax.annotate("", xy=(k2, 0.2), xytext=(k1, 0.2), arrowprops=dict(arrowstyle="-|>", color=RED, lw=1.6))
    ax.text((k1 + k2) / 2, 0.27, "k* falls", color=RED, ha="center", fontsize=8.5)
    ax.set_xticks([k2, k1]); ax.set_xticklabels(["k₂*", "k₁*"]); ax.set_yticks([])
    ax.set_xlabel("Capital per worker k", loc="right")
    ax.legend(frameon=False, fontsize=8.5, loc="upper left")
    save(fig, "solow_breakeven")


# ---------------------------------------------------------------- 6 wars -------
def fig_wars():
    s, d = SOLOW["s1"], SOLOW["delta"]
    ks = kstar(s, d)
    k = np.linspace(0, ks * 1.8, 300)
    fig, axes = plt.subplots(1, 2, figsize=(7.4, 3.2), sharey=True)
    cases = [
        ("War destroys capital\n(Japan, Germany 1945)", ks * 0.4, GREEN, "Δk > 0: fast growth back to k*"),
        ("War kills workers\n(casualties, Pop quiz Q3)", ks * 1.6, RED, "Δk < 0: k and y fall back to k*"),
    ]
    for ax, (title, k1, col, note) in zip(axes, cases):
        ax.plot(k, s * f(k), color=BLUE, lw=2)
        ax.plot(k, d * k, color=ORANGE, lw=2)
        ax.axvline(ks, color=MUTED, lw=0.8, ls=":")
        ax.axvline(k1, color=col, lw=0.8, ls=":")
        ax.scatter([ks], [d * ks], color=INK, s=34, zorder=5)
        ax.scatter([k1, k1], [s * f(k1), d * k1], color=col, s=26, zorder=5)
        ax.annotate("", xy=(ks, 0.08), xytext=(k1, 0.08), arrowprops=dict(arrowstyle="-|>", color=col, lw=1.6))
        ax.set_title(title, fontsize=9.5, fontweight="demibold")
        ax.text(0.5, -0.24, note, transform=ax.transAxes, ha="center", color=col, fontsize=8.5)
        ticks = sorted([(ks, "k*"), (k1, "k₁")])
        ax.set_xticks([t for t, _ in ticks]); ax.set_xticklabels([l for _, l in ticks]); ax.set_yticks([])
        ax.text(k[-1], s * f(k[-1]), " s·f(k)", color=BLUE, fontsize=8, va="center")
        ax.text(k[-1] * 0.8, d * k[-1] * 0.8, "δk ", color=ORANGE, fontsize=8, ha="right")
    save(fig, "wars")


def fig_war_time():
    t = np.linspace(-2, 12, 300)
    fig, axes = plt.subplots(1, 2, figsize=(7.4, 2.5), sharey=True)
    base = 1.0
    paths = [
        ("Capital destroyed", GREEN, lambda x: np.where(x < 0, base, base - 0.45 * np.exp(-0.35 * x))),
        ("Workers killed", RED, lambda x: np.where(x < 0, base, base + 0.35 * np.exp(-0.35 * x))),
    ]
    for ax, (title, col, fn) in zip(axes, paths):
        ax.plot(t, fn(t), color=col, lw=2)
        ax.axhline(base, color=MUTED, lw=0.8, ls=":")
        ax.axvline(0, color=MUTED, lw=0.8)
        ax.text(0.15, 1.42, "war", color=MUTED, fontsize=8)
        ax.set_title(title, fontsize=9.5, fontweight="demibold")
        ax.set_xticks([]); ax.set_yticks([base]); ax.set_yticklabels(["y*"])
        ax.set_xlabel("time", loc="right"); ax.set_ylim(0.45, 1.5)
    axes[0].set_ylabel("output per worker")
    save(fig, "war_time")


# ---------------------------------------------------------------- 7 golden -----
def fig_golden():
    d = SOLOW["delta"]
    a = SOLOW["alpha"]
    kg = (a / d) ** (1 / (1 - a))                  # MPK = a k^(a-1) = d
    k = np.linspace(0, kg * 2.6, 300)
    fig, ax = plt.subplots(figsize=(6.2, 3.6))
    ax.plot(k, f(k), color=INK, lw=2, label="f(k*)  steady-state output")
    ax.plot(k, d * k, color=ORANGE, lw=2, label="δk*  steady-state investment")
    ax.annotate("", xy=(kg, f(kg)), xytext=(kg, d * kg), arrowprops=dict(arrowstyle="<->", color=GREEN, lw=1.6))
    ax.text(kg + 0.8, f(kg) - 0.25, "c* is largest where\nMPK = δ", color=GREEN, fontsize=8.5, va="top", ha="left")
    tang = f(kg) + d * (k - kg)
    ax.plot(k, tang, color=GREEN, lw=0.8, ls="--")
    ax.set_xticks([kg]); ax.set_xticklabels(["k*gold"]); ax.set_yticks([])
    ax.set_xlim(0, k[-1]); ax.set_ylim(0, f(k[-1]) * 1.1)
    ax.text(kg * 0.4, -0.09, "← below gold: raise s", fontsize=8, color=MUTED, ha="center", transform=ax.get_xaxis_transform())
    ax.text(kg * 1.75, -0.09, "above gold: cut s →", fontsize=8, color=MUTED, ha="center", transform=ax.get_xaxis_transform())
    ax.set_xlabel("Steady-state capital per worker k*", loc="right")
    ax.legend(frameon=False, fontsize=8.5, loc="upper left")
    save(fig, "golden")


# ---------------------------------------------------------------- 8 computers --
def fig_computers():
    N, d, n, g = COMPUTERS["N"], COMPUTERS["delta"], COMPUTERS["n"], COMPUTERS["g"]
    left = N * (1 - d)
    trained = N * (1 + g)
    new = N * n
    eff = trained + new
    add = eff - left
    fig, ax = plt.subplots(figsize=(6.2, 2.6))
    rows = [("Effective workers, t+1", [(trained, BLUE, f"{N:g} old × (1 + g) = {trained:g}"), (new, PURPLE, f"+{new:g} new")]),
            ("Computers, t+1", [(left, MUTED, f"{left:g} after depreciation"), (add, GREEN, f"+{add:g} to add")]),
            ("Today", [(N, INK, f"{N:g} computers = {N:g} workers")])]
    for i, (label, parts) in enumerate(rows):
        x = 0
        for w, col, txt in parts:
            ax.barh(i, w, left=x, color=col, height=0.55, alpha=0.85)
            ax.text(x + w / 2, i, txt, color="white", ha="center", va="center", fontsize=8, fontweight="demibold")
            x += w
    ax.set_yticks(range(len(rows))); ax.set_yticklabels([r[0] for r in rows])
    ax.set_xticks([0, N, left, eff]); ax.tick_params(labelsize=8)
    ax.spines["left"].set_visible(False); ax.tick_params(axis="y", length=0)
    save(fig, "computers")


# ---------------------------------------------------------------- surplus ------
def shade(ax, xs, ys, key, text, tx, ty, size=10):
    ax.fill(xs, ys, color=SURPLUS_COLOURS[key], alpha=0.22, lw=0)
    ax.text(tx, ty, text, ha="center", va="center", fontsize=size, fontweight="bold",
            color=SURPLUS_COLOURS[key])


def fig_surplus_general():
    a, b, qr = SURPLUS_GENERAL["demand_top"], SURPLUS_GENERAL["supply_bottom"], SURPLUS_GENERAL["q_restricted"]
    D = lambda q: a - q          # demand
    S = lambda q: b + q          # supply (marginal cost)
    qe = (a - b) / 2; pe = D(qe)
    fig, axes = plt.subplots(1, 2, figsize=(8.4, 3.6))
    for ax, restricted in zip(axes, [False, True]):
        q = np.linspace(0, a, 10)
        ax.plot(q, D(q), color=INK, lw=1.8); ax.plot(q[q <= a - b], S(q[q <= a - b]), color=ORANGE, lw=1.8)
        ax.text(a - 1.2, D(a - 1.2) + 0.5, "Demand", fontsize=9)
        ax.text(7.2, S(7.2) - 1.1, "Supply = MC", fontsize=9, color=ORANGE)
        if not restricted:
            shade(ax, [0, qe, 0], [a, pe, pe], "CS", "CS", qe / 3, (a + 2 * pe) / 3)
            shade(ax, [0, qe, 0], [pe, pe, b], "PS", "PS", qe / 3, (b + 2 * pe) / 3)
            ax.plot([0, qe], [pe, pe], ls="--", color=MUTED, lw=1); ax.plot([qe, qe], [0, pe], ls="--", color=MUTED, lw=1)
            ax.set_xticks([qe]); ax.set_xticklabels(["Q*"]); ax.set_yticks([pe]); ax.set_yticklabels(["P*"])
            ax.set_title("Competitive market: no deadweight loss", fontsize=10)
        else:
            pr = D(qr); mc = S(qr)
            shade(ax, [0, qr, 0], [a, pr, pr], "CS", "CS", qr / 3, (a + 2 * pr) / 3)
            shade(ax, [0, qr, qr, 0], [pr, pr, mc, b], "PS", "PS", qr / 2, (pr + b) / 2 + 0.3)
            shade(ax, [qr, qe, qr], [pr, pe, mc], "DWL", "DWL", qr + (qe - qr) / 3, pe, 9)
            ax.plot([0, qr], [pr, pr], ls="--", color=MUTED, lw=1); ax.plot([qr, qr], [0, pr], ls="--", color=MUTED, lw=1)
            ax.plot([qe, qe], [0, pe], ls=":", color=MUTED, lw=1)
            ax.set_xticks([qr, qe]); ax.set_xticklabels(["Q", "Q*"]); ax.set_yticks([pr]); ax.set_yticklabels(["P"])
            ax.set_title("Output held below Q*: a deadweight loss appears", fontsize=10)
        ax.set_xlim(0, a); ax.set_ylim(0, a + 0.5)
        ax.set_xlabel("Quantity Q", loc="right"); ax.set_ylabel("Price P", loc="top")
    fig.tight_layout()
    save(fig, "surplus_general")


def fig_surplus_table():
    A, C = MARKET["A"], MARKET["C"]
    G = A - C
    fig, axes = plt.subplots(1, 4, figsize=(9.6, 2.9))
    for ax, (name, share, col) in zip(axes, STRUCTURES):
        Q = share * G; P = A - Q
        cs, ps, dwl = Q * (A - P) / 2, (P - C) * Q, (G - Q) * (P - C) / 2
        q = np.linspace(0, A, 10)
        ax.plot(q, A - q, color=INK, lw=1.4); ax.axhline(C, color=ORANGE, lw=1.4)
        shade(ax, [0, Q, 0], [A, P, P], "CS", "CS", Q / 3, (A + 2 * P) / 3 + 4, 9)
        if ps > 0:
            shade(ax, [0, Q, Q, 0], [P, P, C, C], "PS", "PS", Q / 2, (P + C) / 2, 9)
        if dwl > 0:
            shade(ax, [Q, G, Q], [P, C, C], "DWL", "", 0, 0)
            ax.annotate("DWL", (Q + (G - Q) / 3, C + (P - C) / 3), (G + 40, C + 45), fontsize=8.5,
                        fontweight="bold", color=SURPLUS_COLOURS["DWL"], ha="center",
                        arrowprops=dict(arrowstyle="-", color=SURPLUS_COLOURS["DWL"], lw=0.8))
        ax.set_title(name, fontsize=10, color=col, fontweight="bold")
        ax.text(A, A - 5, f"CS {cs:,.0f}\nPS {ps:,.0f}\nDWL {dwl:,.0f}", ha="right", va="top",
                fontsize=7.5, color=MUTED, linespacing=1.3)
        ax.set_xlim(0, A); ax.set_ylim(0, A + 5)
        ax.set_xticks([Q]); ax.set_xticklabels([f"{Q:.0f}"]); ax.set_yticks([C, P]); ax.set_yticklabels([f"{C}", f"{P:.0f}"])
    axes[0].set_ylabel("Price (THB)", loc="top")
    axes[0].text(A, C + 8, "MC", ha="right", color=ORANGE, fontsize=8.5)
    fig.tight_layout()
    save(fig, "surplus_table")


if __name__ == "__main__":
    for fn in [fig_demand, fig_ideas, fig_solow_build, fig_solow_saving, fig_solow_breakeven,
               fig_wars, fig_war_time, fig_golden, fig_computers, fig_surplus_general, fig_surplus_table]:
        fn()
    print("written to", OUT)
