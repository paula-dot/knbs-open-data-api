import { create } from 'zustand'
import axios from 'axios'

export const useStatsStore = create((set) => ({
    populationData: [],
    isLoading: false,
    error: null,

    fetchPopulationData: async () => {
        set({ isLoading: true})
        try {
            // Use relative path so Vite dev server proxy forwards to backend
            const response = await axios.get('/api/v1/data?indicator=POP_TOTAL&year=2019')
            const raw = response.data.data || []

            // Normalize values: handle pgtype.Numeric-like objects ({ String: "123", Valid: true }) or plain numbers/strings
            const normalized = raw.map((d) => {
                const name = d.county_name || d.CountyName || d.county || d.name || d.label || d.code
                const rawVal = d.value ?? d.Value
                let value = 0
                if (rawVal && typeof rawVal === 'object') {
                    // pgtype.Numeric likely has a String field
                    value = Number(rawVal.String ?? rawVal.StringValue ?? 0) || 0
                } else {
                    value = Number(rawVal) || 0
                }
                return { county_name: name, value }
            })

            const sortedData = normalized.sort((a, b) => b.value - a.value)

            set({populationData: sortedData, isLoading: false})
        } catch (err) {
                console.error(err)
                set({ error: 'Failed to fetch population stats', isLoading: false })
            }
        },
    }))