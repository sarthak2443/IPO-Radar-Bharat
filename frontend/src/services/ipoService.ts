import type { IPO, IPOStatus } from '../types'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

export type DataAvailability = 'available' | 'unavailable' | 'estimated' | 'not-announced' | 'delayed'

export type DataSource = {
  name: string
  url?: string
  updatedAt?: string
  availability: DataAvailability
}

export type IPORecord = IPO & {
  source: DataSource
}

export interface IPOService {
  list(): Promise<IPORecord[]>
  getById(id: string): Promise<IPORecord | undefined>
}

const SECTOR_ACCENTS: Record<string, string> = {
  'Automobile': 'linear-gradient(135deg, #3b82f6, #1d4ed8)',
  'Consumer Tech': 'linear-gradient(135deg, #f97316, #ea580c)',
  'Clean Energy': 'linear-gradient(135deg, #10b981, #059669)',
  'Renewable Energy': 'linear-gradient(135deg, #14b8a6, #0d9488)',
  'Financial Services': 'linear-gradient(135deg, #6366f1, #4f46e5)',
  'Electric Vehicles': 'linear-gradient(135deg, #06b6d4, #0891b2)',
  'Engineering': 'linear-gradient(135deg, #8b5cf6, #7c3aed)',
  'Healthcare': 'linear-gradient(135deg, #ec4899, #db2777)',
  'Manufacturing': 'linear-gradient(135deg, #f59e0b, #d97706)',
}

function getInitials(name: string): string {
  if (!name) return 'IP'
  const words = name.replace(/Ltd|Limited|Technologies|India/gi, '').trim().split(/\s+/)
  if (words.length >= 2) {
    return (words[0][0] + words[1][0]).toUpperCase()
  }
  return name.slice(0, 2).toUpperCase()
}

function getSectorAccent(sector: string): string {
  if (SECTOR_ACCENTS[sector]) return SECTOR_ACCENTS[sector]
  return 'linear-gradient(135deg, #10b981, #06b6d4)'
}

function formatDate(dateStr?: string): string {
  if (!dateStr) return 'TBA'
  try {
    const d = new Date(dateStr)
    if (isNaN(d.getTime())) return dateStr
    return d.toLocaleDateString('en-IN', { month: 'short', day: 'numeric', year: 'numeric' })
  } catch {
    return dateStr
  }
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function normalizeIPO(raw: any): IPORecord {
  const companyName = raw.companyName || raw.CompanyName || 'Unknown Company'
  const sector = raw.sector || raw.Sector || 'Diversified'
  const status = (raw.status || raw.Status || 'upcoming').toLowerCase() as IPOStatus

  let priceBand = raw.priceBand || raw.PriceBand || ''
  if (!priceBand) {
    const low = raw.priceBandLow || raw.PriceBandLow || 0
    const high = raw.priceBandHigh || raw.PriceBandHigh || 0
    priceBand = low && high ? `₹${low} – ₹${high}` : 'To be announced'
  }

  let lotSize = raw.lotSize || raw.LotSize || 0
  let minInv = raw.minimumInvestment || raw.MinimumInvestment || ''
  if (!minInv && lotSize && (raw.priceBandHigh || raw.PriceBandHigh)) {
    const price = raw.priceBandHigh || raw.PriceBandHigh
    minInv = `₹${(lotSize * price).toLocaleString('en-IN')}`
  } else if (!minInv) {
    minInv = 'TBA'
  }

  return {
    id: raw.id || raw.ID || companyName.toLowerCase().replace(/\s+/g, '-'),
    companyName,
    shortName: raw.shortName || raw.ShortName || companyName.split(' ')[0],
    category: (raw.category || raw.Category || 'Mainboard') as 'Mainboard' | 'SME',
    sector,
    status,
    priceBand,
    issueSize: raw.issueSize || raw.IssueSize || 'TBA',
    lotSize,
    minimumInvestment: minInv,
    openDate: formatDate(raw.openDate || raw.OpenDate),
    closeDate: formatDate(raw.closeDate || raw.CloseDate),
    listingDate: formatDate(raw.listingDate || raw.ListingDate),
    gmp: Number(raw.gmp ?? raw.GMP ?? 0),
    subscription: Number(raw.subscription ?? raw.Subscription ?? 0),
    radarScore: Number(raw.radarScore ?? raw.RadarScore ?? 75),
    listingGain: raw.listingGain !== undefined ? Number(raw.listingGain) : undefined,
    accent: getSectorAccent(sector),
    initials: getInitials(companyName),
    source: {
      name: raw.exchange || raw.Exchange || 'NSE/BSE Official Feed',
      availability: 'available',
      updatedAt: raw.updatedAt || raw.UpdatedAt,
    },
  }
}

/**
 * Real API Service communicating with the Go Echo backend.
 */
export const ipoService: IPOService = {
  async list(): Promise<IPORecord[]> {
    const response = await fetch(`${API_BASE_URL}/ipos`, {
      headers: { Accept: 'application/json' },
    })
    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`)
    }
    const data = await response.json()
    if (!Array.isArray(data)) {
      return []
    }
    return data.map(normalizeIPO)
  },

  async getById(id: string): Promise<IPORecord | undefined> {
    const response = await fetch(`${API_BASE_URL}/ipos/${encodeURIComponent(id)}`, {
      headers: { Accept: 'application/json' },
    })
    if (response.status === 404) {
      return undefined
    }
    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`)
    }
    const data = await response.json()
    return normalizeIPO(data)
  },
}
