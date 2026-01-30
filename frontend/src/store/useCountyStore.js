import { create } from 'zustand'
import axios from 'axios'

export const useCountyStore = create((set) => ({
  counties: [],
  isLoading: false,
  error: null,
  fetchCounties: async () => {
    set({ isLoading: true, error: null })
    try {
      // Use a relative path so the Vite dev server proxy (vite.config.js) forwards to backend
      const response = await axios.get('/api/v1/counties')
      // Note: Our API returns { "data": [...] }
      set({ counties: response.data.data })
    } catch (err) {
      set({ error: 'Failed to fetch counties' })
    } finally {
      set({ isLoading: false })
    }
  },
}))
