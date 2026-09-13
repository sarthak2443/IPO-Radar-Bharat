import { useEffect, useMemo, useState } from 'react'
import { ipos } from './data/ipos'
import { DetailPanel } from './components/DetailPanel'
import { Icon } from './components/Icon'
import { IPOCard } from './components/IPOCard'
import { Logo } from './components/Logo'
import type { IPO, IPOStatus } from './types'
import './styles.css'

const tabs: { label: string; value: 'all' | IPOStatus }[] = [
  { label: 'All IPOs', value: 'all' }, { label: 'Open now', value: 'open' }, { label: 'Upcoming', value: 'upcoming' }, { label: 'Listed', value: 'listed' },
]

function App() {
  const [activeTab, setActiveTab] = useState<'all' | IPOStatus>('all')
  const [sector, setSector] = useState('All sectors')
  const [query, setQuery] = useState('')
  const [searchOpen, setSearchOpen] = useState(false)
  const [selected, setSelected] = useState<IPO | null>(null)
  const [watched, setWatched] = useState<string[]>(() => JSON.parse(localStorage.getItem('ipo-radar-watchlist') || '[]') as string[])

  useEffect(() => { localStorage.setItem('ipo-radar-watchlist', JSON.stringify(watched)) }, [watched])
  useEffect(() => {
    const handler = (event: KeyboardEvent) => { if ((event.metaKey || event.ctrlKey) && event.key.toLowerCase() === 'k') { event.preventDefault(); setSearchOpen(true) } }
    window.addEventListener('keydown', handler); return () => window.removeEventListener('keydown', handler)
  }, [])

  const toggleWatch = (id: string) => setWatched((items) => items.includes(id) ? items.filter((item) => item !== id) : [...items, id])
  const sectors = ['All sectors', ...new Set(ipos.map((ipo) => ipo.sector))]
  const filtered = useMemo(() => ipos.filter((ipo) => (activeTab === 'all' || ipo.status === activeTab) && (sector === 'All sectors' || ipo.sector === sector) && `${ipo.companyName} ${ipo.sector}`.toLowerCase().includes(query.toLowerCase())), [activeTab, query, sector])
  const openCount = ipos.filter((ipo) => ipo.status === 'open' || ipo.status === 'closing').length
  const avgGain = (ipos.filter((ipo) => ipo.listingGain).reduce((total, ipo) => total + (ipo.listingGain || 0), 0) / ipos.filter((ipo) => ipo.listingGain).length).toFixed(1)

  return <div className="app">
    <header className="topbar"><Logo /><nav><a className="active" href="#discover">Discover</a><a href="#calendar">Calendar</a><a href="#listed">Listed</a><a href="#insights">Insights</a></nav><div className="top-actions"><button className="search-trigger" onClick={() => setSearchOpen(true)}><Icon name="search" size={17} /><span>Search companies</span><kbd>⌘ K</kbd></button><button className="plain-icon" aria-label="Notifications"><Icon name="bell" size={18} /></button><button className="plain-icon" aria-label="Toggle theme"><Icon name="sun" size={18} /></button><button className="profile">SR</button></div></header>
    <main>
      <section className="hero"><div className="eyebrow"><span className="eyebrow-dot" /> MARKET INTELLIGENCE <span className="eyebrow-date">13 SEP 2026</span></div><h1>See the IPO<br /><em>before it hits</em> the market.</h1><p className="hero-copy">A clearer way to discover, research and track India's IPO market — without the noise.</p><div className="hero-actions"><button className="primary-btn" onClick={() => document.getElementById('discover')?.scrollIntoView({ behavior: 'smooth' })}>Explore IPOs <Icon name="arrow" size={16} /></button><button className="secondary-btn" onClick={() => setActiveTab('open')}>See what's open</button></div></section>
      <section className="pulse-section"><div className="section-label"><span>Market pulse</span><span className="live"><i /> Live overview</span></div><div className="pulse-grid"><div className="pulse-lead"><span className="pulse-number">{openCount + 14}</span><span className="pulse-title">IPOs on the radar</span><span className="pulse-caption">Across upcoming, open and listed markets</span></div><div className="pulse-stat"><span>Open now</span><strong>{openCount}</strong><small>+2 since last week</small></div><div className="pulse-stat"><span>Closing soon</span><strong>2</strong><small className="warning">Next 48 hours</small></div><div className="pulse-stat"><span>Avg. listing gain</span><strong className="positive">+{avgGain}%</strong><small>Last 30 days</small></div></div></section>
      <section className="radar-section"><div className="section-heading"><div><span className="eyebrow-small">THE MARKET, AT A GLANCE</span><h2>IPO Radar</h2></div><a href="#calendar">View calendar <Icon name="arrow" size={15} /></a></div><div className="radar-track"><div className="track-line" /><div className="radar-step"><span className="step-count">08</span><span className="step-dot upcoming-dot" /><strong>Upcoming</strong><small>Next 30 days</small></div><div className="radar-step"><span className="step-count">04</span><span className="step-dot open-dot" /><strong>Open now</strong><small>Accepting bids</small></div><div className="radar-step"><span className="step-count">02</span><span className="step-dot closing-dot" /><strong>Closing soon</strong><small>Act in 48 hours</small></div><div className="radar-step"><span className="step-count">06</span><span className="step-dot listed-dot" /><strong>Recently listed</strong><small>Last 30 days</small></div></div></section>
      <section className="discover-section" id="discover"><div className="discover-header"><div><span className="eyebrow-small">CURATED FOR YOU</span><h2>Discover IPOs</h2><p>Make sense of what's moving in the market.</p></div><button className="filter-button"><Icon name="calendar" size={16} /> IPO calendar <Icon name="chevron" size={15} /></button></div><div className="filter-row"><div className="tabs">{tabs.map((tab) => <button className={activeTab === tab.value ? 'active' : ''} key={tab.value} onClick={() => setActiveTab(tab.value)}>{tab.label}</button>)}</div><select value={sector} onChange={(event) => setSector(event.target.value)} aria-label="Filter by sector">{sectors.map((item) => <option key={item}>{item}</option>)}</select></div><div className="cards-grid">{filtered.map((ipo) => <IPOCard key={ipo.id} ipo={ipo} watched={watched.includes(ipo.id)} onWatch={() => toggleWatch(ipo.id)} onSelect={() => setSelected(ipo)} />)}</div>{filtered.length === 0 && <div className="empty">No IPOs match those filters.</div>}</section>
      <section className="watch-section"><div><span className="eyebrow-small">YOUR SHORTLIST</span><h2>My Radar <span>{watched.length}</span></h2><p>Keep the IPOs worth watching close.</p></div>{watched.length ? <div className="watch-chips">{watched.map((id) => { const ipo = ipos.find((item) => item.id === id); return ipo ? <button key={id} onClick={() => setSelected(ipo)}><span style={{ background: ipo.accent }}>{ipo.initials}</span>{ipo.shortName}<Icon name="arrow" size={14} /></button> : null })}</div> : <div className="watch-empty"><Icon name="bookmark" size={20} /><span>Save an IPO to start building your Radar.</span></div>}</section>
      <footer><Logo /><span>Track. Analyze. Decide.</span><span className="footer-right">Data is for informational purposes only · © 2026 IPO Radar</span></footer>
    </main>
    {searchOpen && <div className="search-overlay" onClick={() => setSearchOpen(false)}><div className="search-modal" onClick={(event) => event.stopPropagation()}><div className="search-input"><Icon name="search" size={19} /><input autoFocus value={query} onChange={(event) => setQuery(event.target.value)} placeholder="Search companies, sectors..." /><kbd>ESC</kbd><button onClick={() => setSearchOpen(false)}><Icon name="close" size={16} /></button></div><div className="search-results">{(query ? ipos.filter((ipo) => `${ipo.companyName} ${ipo.sector}`.toLowerCase().includes(query.toLowerCase())) : ipos.slice(0, 4)).map((ipo) => <button key={ipo.id} onClick={() => { setSelected(ipo); setSearchOpen(false) }}><span className="search-mark" style={{ background: ipo.accent }}>{ipo.initials}</span><span><strong>{ipo.companyName}</strong><small>{ipo.category} · {ipo.sector}</small></span><Icon name="arrow" size={15} /></button>)}</div><div className="search-hint">Type to search across 50+ IPOs <span>↵ to select</span></div></div></div>}
    {selected && <DetailPanel ipo={selected} watched={watched.includes(selected.id)} onWatch={() => toggleWatch(selected.id)} onClose={() => setSelected(null)} />}
  </div>
}

export default App
