import type { IPO } from '../types'
import { Icon } from './Icon'

function statusLabel(status: IPO['status']) {
  return { open: 'Open now', closing: 'Closing soon', upcoming: 'Upcoming', listed: 'Listed' }[status]
}

export function IPOCard({ ipo, watched, onWatch, onSelect }: { ipo: IPO; watched: boolean; onWatch: () => void; onSelect: () => void }) {
  return (
    <article className="ipo-card">
      <div className="card-top">
        <div className="company-mark" style={{ background: ipo.accent }}>{ipo.initials}</div>
        <div className="card-company"><h3>{ipo.companyName}</h3><span>{ipo.category} · {ipo.sector}</span></div>
        <button className={`icon-btn ${watched ? 'is-saved' : ''}`} aria-label={watched ? 'Remove from radar' : 'Add to radar'} onClick={onWatch}><Icon name="bookmark" size={17} /></button>
      </div>
      <div className="card-status-row"><span className={`status status-${ipo.status}`}><b />{statusLabel(ipo.status)}</span><span className="card-date">{ipo.status === 'listed' ? `Listed ${ipo.listingDate}` : ipo.status === 'upcoming' ? `Opens ${ipo.openDate}` : `Closes ${ipo.closeDate}`}</span></div>
      <div className="card-price"><strong>{ipo.priceBand}</strong><span>Price band</span></div>
      <div className="card-grid">
        <div><span>Issue size</span><strong>{ipo.issueSize}</strong></div>
        <div><span>Min. investment</span><strong>{ipo.minimumInvestment}</strong></div>
        <div><span>GMP</span><strong className={ipo.gmp >= 0 ? 'positive' : 'negative'}>{ipo.gmp ? `${ipo.gmp > 0 ? '+' : ''}₹${ipo.gmp}` : '—'}</strong></div>
        <div><span>Radar score</span><strong className="score">{ipo.radarScore}<small>/100</small></strong></div>
      </div>
      <button className="text-btn" onClick={onSelect}>View IPO <Icon name="arrow" size={15} /></button>
    </article>
  )
}
