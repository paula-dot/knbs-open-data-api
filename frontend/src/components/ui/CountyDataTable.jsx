import { useState } from 'react'
import {
  Table,
  TableBody,
  TableHeader,
  TableHead,
  TableRow,
  TableCell,
} from '@/components/ui/table'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Search } from 'lucide-react'

export function CountyDataTable({ data = [] }) {
  const [searchTerm, setSearchTerm] = useState('')

  // Simple client-side filtering
  const filteredData = data.filter((county) =>
    county.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    county.code.includes(searchTerm) ||
    (county.former_province && county.former_province.toLowerCase().includes(searchTerm.toLowerCase()))
  )

  return (
    <Card className="w-full shadow-md">
      <CardHeader>
        <CardTitle className="flex justify-between items-center">
          <span>County Registry</span>
          <div className="relative w-64">
            <Search size={16} className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
            <Input
              type="text"
              placeholder="Search counties..."
              className="pl-8"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
        </CardTitle>
      </CardHeader>

      <CardContent>
        <div className="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead className="w-[100px]">Code</TableHead>
                <TableHead>County Name</TableHead>
                <TableHead>Former Province</TableHead>
                <TableHead className="text-right">Area (km²)</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredData.length > 0 ? (
                filteredData.map((county) => (
                  <TableRow key={county.id}>
                    <TableCell className="font-medium">{county.code}</TableCell>
                    <TableCell className="font-bold">{county.name}</TableCell>
                    <TableCell>
                      <Badge variant="secondary" className="font-normal">
                        {county.former_province || 'N/A'}
                      </Badge>
                    </TableCell>
                    <TableCell className="text-right">
                      {/* Format number with commas, handle missing area gracefully */}
                      {county.area_sq_km ? Number(county.area_sq_km).toLocaleString() : '-'}
                    </TableCell>
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell colSpan={4} className="h-24 text-center">
                    No results found.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>

        <div className="text-xs text-muted-foreground mt-4">
          Showing {filteredData.length} of {data.length} counties
        </div>
      </CardContent>
    </Card>
  )
}
