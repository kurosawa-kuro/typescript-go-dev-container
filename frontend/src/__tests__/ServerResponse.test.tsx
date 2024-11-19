import { render, screen } from '@testing-library/react'
import { ServerResponse } from '@/app/components/ServerResponse'

describe('ServerResponse', () => {
  // テストの前後でfetchのモックをセットアップ/クリーンアップ
  beforeEach(() => {
    global.fetch = jest.fn()
  })

  afterEach(() => {
    jest.resetAllMocks()
  })

  // 成功時のモックレスポンス
  const mockSuccessResponse = {
    ok: true,
    json: () => Promise.resolve({ message: 'pong' })
  } as Response

  // エラー時のモックレスポンス
  const mockErrorMessage = 'Failed to fetch data from server'

  it('正常系: サーバーレスポンスが正しく表示される', async () => {
    (global.fetch as jest.Mock).mockResolvedValueOnce(mockSuccessResponse)
    
    render(await ServerResponse())
    
    expect(screen.getByText('Server Response:')).toBeInTheDocument()
    expect(await screen.findByText(/"message":\s*"pong"/)).toBeInTheDocument()
  })

  it('異常系: エラー時に適切なメッセージが表示される', async () => {
    (global.fetch as jest.Mock).mockRejectedValueOnce(new Error('Failed to fetch'))
    
    render(await ServerResponse())
    
    expect(screen.getByText('Server Response:')).toBeInTheDocument()
    expect(await screen.findByText(new RegExp(mockErrorMessage))).toBeInTheDocument()
  })
})