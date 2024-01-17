import type { Metadata } from "next";
import Script from "next/script";
import { Inter } from "next/font/google";
import { Header } from "@/components/common/header/header";
import { Footer } from "@/components/common/footer";
import CssBaseline from "@mui/material/CssBaseline";
import { MetadataDynamic } from "@/components/metadata_dynamic";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "ReMeet kns23",
  description: "神奈川総合高等学校23期生のためのSNS",
  icons: [{ rel: "icon", url: "/favicon.ico" }],
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <html lang='ja'>
        <MetadataDynamic />

        {/* <!-- Google tag (gtag.js) --> */}
        <Script async src='https://www.googletagmanager.com/gtag/js?id=G-LBH4VJGL9S'></Script>
        <Script id='google-analytics'>
          {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', 'G-LBH4VJGL9S');
        `}
        </Script>

        <body className={inter.className} style={{ backgroundColor: "#E0E2EC" }}>
          <CssBaseline />
          <Header />
          {children}
          <Footer />
        </body>
      </html>
    </>
  );
}
