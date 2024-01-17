"use client";

import { useViewport } from "@/hooks/useViewport";

interface Props {}

export function MetadataDynamic({}: Props) {
  const { viewport } = useViewport(); // viewportを取り出す

  return (
    <>
      <meta name='viewport' content={viewport} />
    </>
  );
}
