import type { ReactElement } from 'react'

type IconName = 'search' | 'bell' | 'sun' | 'arrow' | 'bookmark' | 'calendar' | 'chevron' | 'close' | 'spark'

export function Icon({ name, size = 18 }: { name: IconName; size?: number }) {
  const common = { width: size, height: size, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', strokeWidth: 1.8, strokeLinecap: 'round' as const, strokeLinejoin: 'round' as const, 'aria-hidden': true }
  const paths: Record<IconName, ReactElement> = {
    search: <><circle cx="11" cy="11" r="6.5" /><path d="m16 16 4.5 4.5" /></>,
    bell: <><path d="M18 9a6 6 0 0 0-12 0c0 7-3 7-3 9h18c0-2-3-2-3-9M10 22h4" /></>,
    sun: <><circle cx="12" cy="12" r="4" /><path d="M12 2v2m0 16v2M4.93 4.93l1.42 1.42m11.3 11.3 1.42 1.42M2 12h2m16 0h2M4.93 19.07l1.42-1.42m11.3-11.3 1.42-1.42" /></>,
    arrow: <><path d="M5 12h13M13 6l6 6-6 6" /></>,
    bookmark: <path d="M6 4.5A2.5 2.5 0 0 1 8.5 2h7A2.5 2.5 0 0 1 18 4.5V22l-6-3.6L6 22Z" />,
    calendar: <><rect x="3" y="4.5" width="18" height="17" rx="2" /><path d="M16 2v5M8 2v5M3 10h18" /></>,
    chevron: <path d="m7 9 5 5 5-5" />,
    close: <><path d="m6 6 12 12M18 6 6 18" /></>,
    spark: <path d="m12 3 1.4 6.6L20 12l-6.6 1.4L12 20l-1.4-6.6L4 12l6.6-2.4Z" />,
  }
  return <svg {...common}>{paths[name]}</svg>
}
