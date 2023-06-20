import Layout from "@/components/Layout";
import { AppProvider } from "@/contexts";
import { Inter } from "next/font/google";
import { Suspense } from "react";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Expenses",
  description: "Messing with next and go",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <AppProvider>
          <Suspense>
            <Layout>{children}</Layout>
          </Suspense>
        </AppProvider>
      </body>
    </html>
  );
}
