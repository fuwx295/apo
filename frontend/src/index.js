import React from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux'
import 'core-js'

import App from './App'
import { store } from './store/store'
import { ToastProvider } from './components/Toast/ToastContext'
import { ConfigProvider, theme } from 'antd'

import posthog from 'posthog-js'
import { PostHogProvider } from 'posthog-js/react'
import { MessageProvider } from './contexts/MessageContext'
import ErrorBoundary from './components/ErrorBoundary'
import zhCN from 'antd/es/locale/zh_CN' // 引入中文包
const apiHost = import.meta.env.VITE_PUBLIC_POSTHOG_HOST
const apiKey = import.meta.env.VITE_PUBLIC_POSTHOG_KEY

posthog.init(apiKey, {
  api_host: apiHost,
  person_profiles: 'identified_only',
})
const AppWrapper = () => {
  return (
    <ErrorBoundary>
      <Provider store={store}>
        <ToastProvider>
          <ConfigProvider
            locale={zhCN}
            theme={{
              algorithm: theme.darkAlgorithm,
              components: {
                Segmented: {
                  itemSelectedBg: '#4096ff',
                },
                Layout: {
                  bodyBg: '#1d222b',
                  siderBg: '#1d222b',
                },
              },
            }}
          >
            <MessageProvider>
              <App />
            </MessageProvider>
          </ConfigProvider>
        </ToastProvider>
      </Provider>
    </ErrorBoundary>
  )
}
createRoot(document.getElementById('root')).render(
  apiKey && apiHost ? (
    <PostHogProvider client={posthog}>
      <AppWrapper />
    </PostHogProvider>
  ) : (
    <AppWrapper />
  ),
)
