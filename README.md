# grafana-server
Grafana JSON Server in Go


    / should return 200 ok. Used for "Test connection" on the datasource config page.
    /search used by the find metric options on the query tab in panels.
    /query should return metrics based on input.
    /annotations should return annotations.
    /tag-keys should return tag keys for ad hoc filters.
    /tag-values should return tag values for ad hoc filters.
