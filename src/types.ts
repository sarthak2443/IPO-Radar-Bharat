export type IPOStatus = 'upcoming' | 'open' | 'closing' | 'listed'

export type IPO = {
  id: string
  companyName: string
  shortName: string
  category: 'Mainboard' | 'SME'
  sector: string
  status: IPOStatus
  priceBand: string
  issueSize: string
  lotSize: number
  minimumInvestment: string
  openDate: string
  closeDate: string
  listingDate: string
  gmp: number
  subscription: number
  radarScore: number
  listingGain?: number
  accent: string
  initials: string
}
