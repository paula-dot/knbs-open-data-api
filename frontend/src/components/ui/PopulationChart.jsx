import React from 'react'
import {
  ResponsiveContainer,
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
} from 'recharts'

function numberFormatter(val) {
  if (val === null || val === undefined) return '-'
  // Ensure number and format with commas
  const num = Number(val)
  if (Number.isNaN(num)) return String(val)
  return num.toLocaleString()
}

const CustomTooltip = ({ active, payload, label }) => {
  if (!active || !payload || !payload.length) return null
  const value = payload[0].value
  return (
    <div className="bg-white p-2 rounded shadow border">
      <div className="text-sm font-medium">{label}</div>
      <div className="text-xs text-muted-foreground">{numberFormatter(value)}</div>
    </div>
  )
}

export function PopulationChart({ data = [], topN = 10 }) {
  // Expect data items like: { county_name: 'Nairobi', value: 12345 } or { county: 'Nairobi', value: '4397073' }
  const prepared = data
    .map((d) => {
      const name = d.county_name || d.CountyName || d.county || d.name || d.label || d.code
      const rawVal = d.value ?? d.Value
      const value = Number(rawVal) || 0
      return { name, value }
    })
    .filter((d) => d.name)
    .slice(0, topN)

  return (
    <div className="w-full h-80 bg-white rounded-md shadow p-4">
      <h3 className="text-lg font-semibold mb-2">Population (2019)</h3>
      <ResponsiveContainer width="100%" height="90%">
        <BarChart data={prepared} layout="vertical" margin={{ top: 10, right: 20, bottom: 10, left: 50 }}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis type="number" tickFormatter={(v) => numberFormatter(v)} />
          <YAxis type="category" dataKey="name" width={150} />
          <Tooltip content={<CustomTooltip />} formatter={(val) => numberFormatter(val)} />
          <Bar dataKey="value" fill="#2563EB" radius={[4, 4, 4, 4]} />
        </BarChart>
      </ResponsiveContainer>
    </div>
  )
}
