import type { Metadata } from "next";
import { Noto_Sans_JP } from "next/font/google";
import "@/app/globals.css";

// Noto Sans JPフォントを設定
const notoSansJP = Noto_Sans_JP({
  subsets: ['latin'],
  weight: ['400', '700'],  // 必要な太さを指定
  preload: true,
});

export const metadata: Metadata = {
  title: "Go Node.js RDB App",
  description: "Go Node.js RDB App",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja" className="dark" suppressHydrationWarning>
      <body className={`${notoSansJP.className} antialiased`} suppressHydrationWarning>
        {children}
      </body>
    </html>
  );
}
