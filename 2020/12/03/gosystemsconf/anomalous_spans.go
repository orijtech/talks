                // Return anomalous spans.
                µPlus3σ := st.mean + 3*st.stddev
                µMinus3σ := st.mean - 3*st.stddev
                above3Sigma := make([]*spanDetails, 0, len(sds))
                below3Sigma := make([]*spanDetails, 0, len(sds))
                nonInteresting := make([]*spanDetails, 0, len(sds))
                for _, sd := range sds {
                        if sd.latencyMs >= µPlus3σ {
                                above3Sigma = append(above3Sigma, sd)
                        } else if sd.latencyMs <= µMinus3σ {
                                below3Sigma = append(below3Sigma, sd)
                        } else {
                                nonInteresting = append(nonInteresting, sd)
                        }
                }