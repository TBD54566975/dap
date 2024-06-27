export async function GET() {
  return new Response(
    JSON.stringify({
      "id": "did:web:daps.dev",
      "service": [
        {
          "id": "DAPRegistry",
          "type": "DAPRegistry",
          "serviceEndpoint": [
            "https://didpay.me"
          ]
        }
      ]
    })
  )
}