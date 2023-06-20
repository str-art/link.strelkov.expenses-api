import React from "react";
import { NextAppDirEmotionCacheProvider } from "tss-react/next/appDir";
import ThemeProvider from "./ThemeProvider";
import { NextThemeProvider } from "./NextThemeProvider";
import ClientsProvider from "./ClientProvider";

type AppProviderProps = {
  children: React.ReactNode;
};

const AppProvider = (props: AppProviderProps) => {
  const { children } = props;
  return (
    <NextThemeProvider>
      <NextAppDirEmotionCacheProvider options={{ key: "css" }}>
        <ThemeProvider>
          <ClientsProvider apiUrl={process.env.API_URL || "api"}>
            {children}
          </ClientsProvider>
        </ThemeProvider>
      </NextAppDirEmotionCacheProvider>
    </NextThemeProvider>
  );
};

export default AppProvider;
