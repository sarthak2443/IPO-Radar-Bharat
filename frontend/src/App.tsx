import { useEffect, useMemo, useState } from 'react'
import { ipoService, type IPORecord } from './services/ipoService'
import { DetailPanel } from './components/DetailPanel'
import { Icon } from './components/Icon'
import { IPOCard } from './components/IPOCard'
import { Logo } from './components/Logo'
import type { IPOStatus } from './types'
import './styles.css'

const tabs: { label: string; value: 'all' | IPOStatus }[] = [
  { label: 'All IPOs', value: 'all' },
  { label: 'Open now', value: 'open' },
  { label: 'Upcoming', value: 'upcoming' },
  { label: 'Listed', value: 'listed' },
]

function App() {
  const [darkMode, setDarkMode] = useState(() => {
    const savedTheme = localStorage.getItem('ipo-radar-theme')
    return savedTheme ? savedTheme === 'dark' : true // Default to better dark mode
  })

  const [ipos, setIpos] = useState<IPORecord[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const [activeTab, setActiveTab] = useState<'all' | IPOStatus>('all')
  const [sector, setSector] = useState('All sectors')
  const [query, setQuery] = useState('')
  const [searchOpen, setSearchOpen] = useState(false)
  const [selected, setSelected] = useState<IPORecord | null>(null)
  const [watched, setWatched] = useState<string[]>(() => {
    try {
      return JSON.parse(localStorage.getItem('ipo-radar-watchlist') || '[]') as string[]
    } catch {
      return []
    }
  })

  const fetchIPOs = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await ipoService.list()
      setIpos(data)
    } catch (err) {
      setError((err as Error).message || 'Failed to reach API server')
      setIpos([])
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchIPOs()
  }, [])

  useEffect(() => {
    localStorage.setItem('ipo-radar-watchlist', JSON.stringify(watched))
  }, [watched])

  useEffect(() => {
    document.documentElement.dataset.theme = darkMode ? 'dark' : 'light'
    localStorage.setItem('ipo-radar-theme', darkMode ? 'dark' : 'light')
  }, [darkMode])

  useEffect(() => {
    const handler = (event: KeyboardEvent) => {
      if ((event.metaKey || event.ctrlKey) && event.key.toLowerCase() === 'k') {
        event.preventDefault()
        setSearchOpen(true)
      }
    }
    window.addEventListener('keydown', handler)
    return () => window.removeEventListener('keydown', handler)
  }, [])

  const toggleWatch = (id: string) => {
    setWatched((items) =>
      items.includes(id) ? items.filter((item) => item !== id) : [...items, id]
    )
  }

  const sectors = useMemo(() => {
    return ['All sectors', ...new Set(ipos.map((ipo) => ipo.sector).filter(Boolean))]
  }, [ipos])

  const filtered = useMemo(() => {
    return ipos.filter((ipo) => {
      const matchesTab = activeTab === 'all' || ipo.status === activeTab
      const matchesSector = sector === 'All sectors' || ipo.sector === sector
      const matchesQuery =
        !query ||
        `${ipo.companyName} ${ipo.sector} ${ipo.shortName}`
          .toLowerCase()
          .includes(query.toLowerCase())
      return matchesTab && matchesSector && matchesQuery
    })
  }, [ipos, activeTab, sector, query])

  const openCount = ipos.filter((ipo) => ipo.status === 'open' || ipo.status === 'closing').length
  const upcomingCount = ipos.filter((ipo) => ipo.status === 'upcoming').length
  const listedCount = ipos.filter((ipo) => ipo.status === 'listed').length

  const avgGain = useMemo(() => {
    const listedWithGain = ipos.filter((ipo) => ipo.listingGain !== undefined)
    if (listedWithGain.length === 0) return '0.0'
    const sum = listedWithGain.reduce((total, ipo) => total + (ipo.listingGain || 0), 0)
    return (sum / listedWithGain.length).toFixed(1)
  }, [ipos])

  return (
    <div className="app">
      <header className="topbar">
        <Logo />
        <nav>
          <a className="active" href="#discover">Discover</a>
          <a href="#pulse">Market Pulse</a>
          <a href="#radar">Radar</a>
          <a href="#watchlist">Watchlist</a>
        </nav>
        <div className="top-actions">
          <button className="search-trigger" onClick={() => setSearchOpen(true)}>
            <Icon name="search" size={17} />
            <span>Search companies</span>
            <kbd>⌘ K</kbd>
          </button>
          <button
            className="plain-icon"
            aria-label={darkMode ? 'Switch to light theme' : 'Switch to dark theme'}
            onClick={() => setDarkMode((v) => !v)}
            title={darkMode ? 'Switch to light theme' : 'Switch to dark theme'}
          >
            <Icon name="sun" size={18} />
          </button>
          <button className="profile" title="IPO Radar Bharat">IN</button>
        </div>
      </header>

      <main>
        {error && (
          <div className="api-banner error-banner">
            <div className="banner-content">
              <strong>Backend Offline:</strong> {error}
              <span>Make sure the backend is running at <code>http://localhost:8080</code></span>
            </div>
            <button className="banner-btn" onClick={fetchIPOs}>Retry Connection</button>
          </div>
        )}

        <section className="hero">
          <div className="eyebrow">
            <span className="eyebrow-dot" />
            LIVE MARKET INTELLIGENCE
            <span className="eyebrow-date">LIVE BACKEND SYNC</span>
          </div>
          <h1>
            Track Indian IPOs<br />
            <em>with precision</em> and clarity.
          </h1>
          <p className="hero-copy">
            Institutional-grade IPO discovery, live subscription demand, and verified exchange listings.
          </p>
          <div className="hero-actions">
            <button
              className="primary-btn"
              onClick={() => document.getElementById('discover')?.scrollIntoView({ behavior: 'smooth' })}
            >
              Explore IPOs <Icon name="arrow" size={16} />
            </button>
            <button className="secondary-btn" onClick={fetchIPOs}>
              Refresh Feed
            </button>
          </div>
        </section>

        <section className="pulse-section" id="pulse">
          <div className="section-label">
            <span>Market pulse</span>
            <span className="live"><i /> Real-time status</span>
          </div>
          <div className="pulse-grid">
            <div className="pulse-lead">
              <span className="pulse-number">{ipos.length}</span>
              <span className="pulse-title">Total tracked IPOs</span>
              <span className="pulse-caption">Synced from PostgreSQL database</span>
            </div>
            <div className="pulse-stat">
              <span>Open now</span>
              <strong>{openCount}</strong>
              <small>Accepting bids</small>
            </div>
            <div className="pulse-stat">
              <span>Upcoming</span>
              <strong>{upcomingCount}</strong>
              <small>Filing / Pre-issue</small>
            </div>
            <div className="pulse-stat">
              <span>Avg. listing gain</span>
              <strong className={Number(avgGain) >= 0 ? 'positive' : 'negative'}>
                {Number(avgGain) >= 0 ? '+' : ''}{avgGain}%
              </strong>
              <small>{listedCount} listed issues</small>
            </div>
          </div>
        </section>

        <section className="radar-section" id="radar">
          <div className="section-heading">
            <div>
              <span className="eyebrow-small">MARKET STAGES</span>
              <h2>IPO Lifecycle Radar</h2>
            </div>
            <span className="live-tag">Active Tracker</span>
          </div>
          <div className="radar-track">
            <div className="track-line" />
            <div className="radar-step">
              <span className="step-count">01</span>
              <span className="step-dot upcoming-dot" />
              <strong>Upcoming</strong>
              <small>{upcomingCount} issues</small>
            </div>
            <div className="radar-step">
              <span className="step-count">02</span>
              <span className="step-dot open-dot" />
              <strong>Open Now</strong>
              <small>{openCount} accepting bids</small>
            </div>
            <div className="radar-step">
              <span className="step-count">03</span>
              <span className="step-dot closing-dot" />
              <strong>Closing</strong>
              <small>Final allotment phase</small>
            </div>
            <div className="radar-step">
              <span className="step-count">04</span>
              <span className="step-dot listed-dot" />
              <strong>Listed</strong>
              <small>{listedCount} on exchange</small>
            </div>
          </div>
        </section>

        <section className="discover-section" id="discover">
          <div className="discover-header">
            <div>
              <span className="eyebrow-small">OFFICIAL FEEDS</span>
              <h2>Discover Issues</h2>
              <p>Filter by issue stage and market industry.</p>
            </div>
          </div>

          <div className="filter-row">
            <div className="tabs">
              {tabs.map((tab) => (
                <button
                  className={activeTab === tab.value ? 'active' : ''}
                  key={tab.value}
                  onClick={() => setActiveTab(tab.value)}
                >
                  {tab.label}
                </button>
              ))}
            </div>
            <select
              value={sector}
              onChange={(e) => setSector(e.target.value)}
              aria-label="Filter by sector"
            >
              {sectors.map((item) => (
                <option key={item} value={item}>{item}</option>
              ))}
            </select>
          </div>

          {loading ? (
            <div className="cards-grid">
              {[1, 2, 3].map((i) => (
                <div key={i} className="ipo-card skeleton-card">
                  <div className="skeleton-line title" />
                  <div className="skeleton-line" />
                  <div className="skeleton-line short" />
                </div>
              ))}
            </div>
          ) : filtered.length > 0 ? (
            <div className="cards-grid">
              {filtered.map((ipo) => (
                <IPOCard
                  key={ipo.id}
                  ipo={ipo}
                  watched={watched.includes(ipo.id)}
                  onWatch={() => toggleWatch(ipo.id)}
                  onSelect={() => setSelected(ipo)}
                />
              ))}
            </div>
          ) : (
            <div className="empty">
              {ipos.length === 0 ? (
                <div>
                  <h3>No IPOs in database yet</h3>
                  <p>
                    Start the backend and run <code>make seed</code> in <code>backend/</code> to populate live Indian IPOs into PostgreSQL.
                  </p>
                  <button className="primary-btn" onClick={fetchIPOs} style={{ margin: '14px auto 0' }}>
                    Reload Database
                  </button>
                </div>
              ) : (
                <p>No IPOs match the selected filter criteria.</p>
              )}
            </div>
          )}
        </section>

        <section className="watch-section" id="watchlist">
          <div>
            <span className="eyebrow-small">YOUR RADAR</span>
            <h2>Watchlist <span>{watched.length}</span></h2>
            <p>Pinned IPOs for quick subscription tracking.</p>
          </div>
          {watched.length ? (
            <div className="watch-chips">
              {watched.map((id) => {
                const ipo = ipos.find((item) => item.id === id)
                return ipo ? (
                  <button key={id} onClick={() => setSelected(ipo)}>
                    <span style={{ background: ipo.accent }}>{ipo.initials}</span>
                    {ipo.shortName}
                    <Icon name="arrow" size={14} />
                  </button>
                ) : null
              })}
            </div>
          ) : (
            <div className="watch-empty">
              <Icon name="bookmark" size={20} />
              <span>Click the bookmark icon on any IPO card to track it here.</span>
            </div>
          )}
        </section>

        <footer>
          <Logo />
          <span>Real-time IPO Intelligence & Analytics Platform</span>
          <span className="footer-right">
            Connected to Go API · Data for educational & analytical purposes
          </span>
        </footer>
      </main>

      {searchOpen && (
        <div className="search-overlay" onClick={() => setSearchOpen(false)}>
          <div className="search-modal" onClick={(e) => e.stopPropagation()}>
            <div className="search-input">
              <Icon name="search" size={19} />
              <input
                autoFocus
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                placeholder="Search companies, sectors..."
              />
              <kbd>ESC</kbd>
              <button onClick={() => setSearchOpen(false)}>
                <Icon name="close" size={16} />
              </button>
            </div>
            <div className="search-results">
              {(query
                ? ipos.filter((ipo) =>
                    `${ipo.companyName} ${ipo.sector}`
                      .toLowerCase()
                      .includes(query.toLowerCase())
                  )
                : ipos.slice(0, 5)
              ).map((ipo) => (
                <button
                  key={ipo.id}
                  onClick={() => {
                    setSelected(ipo)
                    setSearchOpen(false)
                  }}
                >
                  <span className="search-mark" style={{ background: ipo.accent }}>
                    {ipo.initials}
                  </span>
                  <span>
                    <strong>{ipo.companyName}</strong>
                    <small>{ipo.category} · {ipo.sector}</small>
                  </span>
                  <Icon name="arrow" size={15} />
                </button>
              ))}
            </div>
            <div className="search-hint">
              Search by company name or sector <span>↵ to view</span>
            </div>
          </div>
        </div>
      )}

      {selected && (
        <DetailPanel
          ipo={selected}
          watched={watched.includes(selected.id)}
          onWatch={() => toggleWatch(selected.id)}
          onClose={() => setSelected(null)}
        />
      )}
    </div>
  )
}

export default App
