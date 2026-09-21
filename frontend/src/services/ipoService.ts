import { ipos } from '../data/ipos'
import type { IPO } from '../types'

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

/**
 * The UI talks to this interface rather than to a provider or scraper.
 * Replace this adapter with a server API when live ingestion is enabled.
 */
export const mockIPOService: IPOService = {
  async list() {
    return ipos.map((ipo) => ({
      ...ipo,
      source: {
        name: 'IPO Radar development dataset',
        availability: 'estimated',
      },
    }))
  },
  async getById(id) {
    const ipo = ipos.find((item) => item.id === id)
    return ipo
      ? {
          ...ipo,
          source: {
            name: 'IPO Radar development dataset',
            availability: 'estimated',
          },
        }
      : undefined
  },
}
