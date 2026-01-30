import { create } from 'zustand'
import axios from 'axios'

interface County {
    id: number,
    name: string,
    code: string,
    former_province: string | null
}

interface CountyStore {
    counties: County[]
    isLoading: boolean
    error: string | null
    fetchCounties: () => Promise<void>
}

export const useCountyStore = create<CountyStore>((set) => ({
    counties: [],
    isLoading: false,
    error: null,
    fetchCounties: async () => {
        set({ isLoading: true, error: null })
        try {
            const response = await axios.get<{ data: County[] }>('/api/v1/counties')
            set({ counties: response.data.data })
        } catch (err: any) {
            set({ error: err?.message ?? String(err) })
        } finally {
            set({ isLoading: false })
        }
    },
}))