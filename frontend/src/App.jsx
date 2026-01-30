import { useEffect } from 'react'
import { useCountyStore } from './store/useCountyStore'
import { useStatsStore } from './store/useStatsStore' // <--- Import
import { CountyDataTable } from './components/ui/CountyDataTable'
import { PopulationChart } from './components/ui/PopulationChart' // <--- Import

function App() {
  const { counties, isLoading, error, fetchCounties } = useCountyStore()
  // Destructure stats store
  const { populationData, fetchPopulationData } = useStatsStore()

  useEffect(() => {
    fetchCounties()
    fetchPopulationData()
  }, [fetchCounties, fetchPopulationData])

  if (isLoading) return <div className="p-10">Loading Kenya's data...</div>
  if (error) return <div className="p-10 text-red-500">{error}</div>

  return (
    <div className="min-h-screen bg-gray-50/50 p-8 font-sans text-gray-900">
      <div className="max-w-5xl mx-auto space-y-6">
        <div className="space-y-2">
          <h1 className="text-4xl font-extrabold tracking-tight">🇰🇪 Kenya Open Data</h1>
          <p className="text-lg text-muted-foreground">
            Official standardized metadata and statistics.
          </p>
        </div>

        {/* --- NEW CHART SECTION --- */}
        {populationData && populationData.length > 0 && (
          <div className="grid grid-cols-1 gap-4">
             <PopulationChart data={populationData} />
          </div>
        )}

        {/* Existing County Table */}
        <CountyDataTable data={counties} />
      </div>
    </div>
  )
}

export default App
