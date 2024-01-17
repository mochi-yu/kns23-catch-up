'use client'

import { useLayoutEffect, useState } from 'react'

const standardPC = 1000
const standardSP = 768
const minimumSP = 320

function getViewport(width: number): string {
  if (width > standardPC) {
    return 'device-width'
  } else if (width > standardSP) {
    return '1200'
  } else if (width >= minimumSP) {
    return 'device-width'
  } else {
    return '320'
  }
}

export const useViewport = () => {
  const [viewport, setViewport] = useState('width=device-width,initial-scale=1.0,maximum-scale=1.0')

  const handleWindowResize = () => {
    if (window != undefined) {
      setViewport(`width=${getViewport(window.outerWidth)},maximum-scale=1.0`)
    } else {
      return
    }
  }

  useLayoutEffect(() => {
    handleWindowResize()
    window.addEventListener('resize', handleWindowResize)
    return () => window.removeEventListener('resize', handleWindowResize)
  }, [])

  return { viewport }
}