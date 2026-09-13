import type { CSSProperties } from 'react'
import type { IPO } from '../types'
import { Icon } from './Icon'

export function DetailPanel({ ipo, watched, onWatch, onClose }: { ipo: IPO; watched: boolean; onWatch: () => void; onClose: () => void }) {
  return <div className="detail-overlay" role="dialog" aria-modal="true" aria-label={`${ipo.companyName} details`} onClick={onClose}>
    <aside className="detail-panel" onClick={(event) => event.stopPropagation()}>
      <div className="detail-header"><button className="back-btn" onClick={onClose}><Icon name="close" size={18} /> Close</button><button className={`icon-btn ${watched ? 'is-saved' : ''}`} onClick={onWatch}><Icon name="bookmark" size={17} /></button></div>
      <div className="detail-identity"><div className="company-mark large" style={{ background: ipo.accent }}>{ipo.initials}</div><div><span className={`status status-${ipo.status}`}><b />{ipo.status === 'closing' ? 'Closing soon' : ipo.status === 'open' ? 'Open now' : ipo.status === 'listed' ? 'Listed' : 'Upcoming'}</span><h2>{ipo.companyName}</h2><p>{ipo.category} · {ipo.sector}</p></div></div>
      <div className="detail-score"><div><span>IPO Radar score</span><strong>{ipo.radarScore}<small>/100</small></strong></div><div className="score-ring" style={{ '--score': `${ipo.radarScore * 3.6}deg` } as CSSProperties}><span>{ipo.radarScore}</span></div></div>
      <div className="detail-key"><div><span>Price band</span><strong>{ipo.priceBand}</strong></div><div><span>Issue size</span><strong>{ipo.issueSize}</strong></div><div><span>Lot size</span><strong>{ipo.lotSize} shares</strong></div><div><span>Min. investment</span><strong>{ipo.minimumInvestment}</strong></div></div>
      <button className={`primary-btn wide ${watched ? 'saved' : ''}`} onClick={onWatch}><Icon name="bookmark" size={16} /> {watched ? 'Added to your Radar' : 'Add to my Radar'}</button>
      <section className="detail-section"><div className="section-heading"><h3>IPO timeline</h3><span>2026</span></div><div className="mini-timeline"><div className="done"><b />RHP filed<small>Aug 28</small></div><div className="done"><b />Opens<small>{ipo.openDate.replace(', 2026', '')}</small></div><div className="active"><b />Closes<small>{ipo.closeDate.replace(', 2026', '')}</small></div><div><b />Listing<small>{ipo.listingDate.replace(', 2026', '')}</small></div></div></section>
      <section className="detail-section"><div className="section-heading"><h3>Demand signals</h3><span>As of today</span></div><div className="signal"><div><span>Subscription</span><strong>{ipo.subscription ? `${ipo.subscription}×` : 'Not open'}</strong></div><div className="signal-bar"><i style={{ width: `${Math.min(ipo.subscription * 4, 100)}%` }} /></div></div><div className="signal"><div><span>GMP</span><strong className={ipo.gmp >= 0 ? 'positive' : 'negative'}>{ipo.gmp ? `₹${ipo.gmp}` : 'Not available'}</strong></div><div className="gmp-note">Unofficial market indicator</div></div></section>
      <p className="detail-disclaimer">IPO Radar score and GMP are analytical indicators, not investment advice. Market data may be delayed.</p>
    </aside>
  </div>
}
