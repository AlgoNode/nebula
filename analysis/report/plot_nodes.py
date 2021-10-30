from collections import OrderedDict

import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
from matplotlib import ticker
from lib_agent import agent_name, known_agents, go_ipfs_version_mappings, go_ipfs_version
from lib_db import DBClient, calendar_week
from lib_fmt import fmt_thousands, fmt_barplot, thousands_ticker_formatter, fmt_percentage

sns.set_theme()

client = DBClient()

all_peer_ids = set(client.get_all_peer_ids())
online_peer_ids = set(client.get_online_peer_ids())
offline_peer_ids = set(client.get_offline_peer_ids())
all_entering_peer_ids = set(client.get_entering_peer_ids())
all_leaving_peer_ids = set(client.get_leaving_peer_ids())
dangling_peer_ids = set(client.get_dangling_peer_ids())
one_off_peer_ids = set(client.get_oneoff_peer_ids())

entering_peer_ids = all_entering_peer_ids.difference(all_leaving_peer_ids)
leaving_peer_ids = all_leaving_peer_ids.difference(all_entering_peer_ids)

assert online_peer_ids.issubset(all_peer_ids)
assert offline_peer_ids.issubset(all_peer_ids)
assert entering_peer_ids.issubset(all_peer_ids)
assert leaving_peer_ids.issubset(all_peer_ids)
assert dangling_peer_ids.issubset(all_peer_ids)

assert online_peer_ids.isdisjoint(offline_peer_ids)
assert online_peer_ids.isdisjoint(entering_peer_ids)
assert online_peer_ids.isdisjoint(leaving_peer_ids)
assert online_peer_ids.isdisjoint(dangling_peer_ids)
assert offline_peer_ids.isdisjoint(entering_peer_ids)

assert offline_peer_ids.isdisjoint(leaving_peer_ids)
assert offline_peer_ids.isdisjoint(dangling_peer_ids)

assert entering_peer_ids.isdisjoint(leaving_peer_ids)
assert entering_peer_ids.isdisjoint(dangling_peer_ids)

assert leaving_peer_ids.isdisjoint(dangling_peer_ids)

data = OrderedDict([
    ("online", len(online_peer_ids)),
    ("offline", len(offline_peer_ids)),
    ("entered", len(entering_peer_ids)),
    ("left", len(leaving_peer_ids)),
    ("dangling", len(dangling_peer_ids)),
])
data = OrderedDict(reversed(sorted(data.items(), key=lambda item: item[1])))

# Plotting

fig, (ax) = plt.subplots(figsize=(8, 5))  # rows, cols

sns.barplot(ax=ax, x=list(data.keys()), y=list(data.values()))
ax.get_yaxis().set_major_formatter(thousands_ticker_formatter)
ax.bar_label(ax.containers[0], list(map(fmt_percentage(len(all_peer_ids)), data.values())))

ax.title.set_text(f"Node Classification of {fmt_thousands(len(all_peer_ids))} Visited Peers")
ax.set_ylabel("Count")

plt.tight_layout()

plt.savefig(f"./plots-{calendar_week}/nodes.png")
plt.show()