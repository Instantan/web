import { Button } from "./Button"
import { PackageOpen, Zap, FileJson, Rocket, Unplug, FileType2, Terminal, Webhook, Copy } from "lucide-react"
import { Gopher } from "./Gopher"
import './index.css'

export default function LandingPage() {
  return (
    <div className="flex flex-col min-h-screen bg-default text-gray-100">
      <header className="px-4 lg:px-6 h-14 flex items-center border-b border-gray-800">
        <a className="flex items-center justify-center" href="#">
          <Webhook className="h-6 w-6 mr-2 text-emerald-400" />
          <span className="font-bold text-lg text-emerald-400">web</span>
        </a>
        <nav className="ml-auto flex gap-4 sm:gap-6">
          <a className="text-sm font-medium hover:text-emerald-400 transition-colors" href="#docs">
            Docs
          </a>
          <a className="text-sm font-medium hover:text-emerald-400 transition-colors" href="https://github.com/Instantan/web">
            GitHub
          </a>
        </nav>
      </header>
      <main className="flex-1">
        <section style={{ paddingTop: 120, paddingBottom: 120 }} className="w-full py-8 md:py-20 lg:py-28 xl:py-44 relative overflow-hidden">
          <div className="absolute inset-0 bg-grid-white/[0.02] bg-[size:50px_50px]" />
          <div className="container px-4 md:px-6 relative z-10">
            <div className="flex flex-col items-center text-center">
              <div>
                <h1 style={{ fontWeight: 900, lineHeight: 1.15 }} className="text-4xl font-bold tracking-tighter sm:text-5xl md:text-6xl lg:text-7xl/none bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">
                  The Zero-Dependency<br />Go Framework
                </h1>
                <p className="mx-auto my-7 max-w-[700px] text-gray-400 md:text-xl">
                  Build robust web applications with pure Go. No external dependencies, just the power of the standard library.
                </p>
              </div>
              <div className="mt-2">
                <Button className="bg-emerald-500 hover:bg-emerald-600 text-gray-950 px-10">Get Started</Button>
              </div>
              <Gopher />
            </div>
          </div>
          <div className="absolute inset-0 bg-gradient-to-t from-gray-950 to-transparent pointer-events-none" />
        </section>
        <section className="w-full py-12 md:py-24 lg:py-32 bg-gray-900">
          <div className="container px-4 md:px-6">
            <div className="flex flex-col items-center space-y-4 text-center">
              <Terminal className="h-12 w-12 text-emerald-400 mb-4" />
              <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">
                Installation
              </h2>
              <p className="mx-auto max-w-[700px] text-gray-400 md:text-xl mb-4">
                Get started with web in just one command. It's as simple as that!
              </p>
              <div className="relative w-full max-w-2xl bg-gray-800 rounded-lg overflow-hidden transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(20,255,180,0.3)]">
                <pre className="language-bash">
                  <code className="text-sm text-gray-300 p-4 block overflow-x-auto">
                    go get github.com/Instantan/web
                    <button
                      className="absolute top-2 right-2 p-1 rounded-md bg-gray-700 hover:bg-gray-600 transition-colors"
                    >
                      <Copy className="h-4 w-4 text-gray-400" />
                    </button>
                  </code>
                </pre>
              </div>
            </div>
          </div>
        </section>
        <section id="features" className="w-full py-12 md:py-24 lg:py-32 bg-default">
          <div className="container px-4 md:px-6">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-8 bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">Features</h2>
            <div className="grid gap-10 sm:grid-cols-2 md:grid-cols-3">
              <div className="flex flex-col items-center space-y-2 p-6 rounded-lg bg-gray-800 transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(52,211,153,0.3)]">
                <div className="relative p-2 rounded-lg">
                  <div className="absolute inset-0 rounded-lg bg-emerald-400/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300 blur-md"></div>
                  <PackageOpen className="h-8 w-8 text-emerald-400 relative z-10" />
                </div>
                <h3 className="text-xl font-bold text-emerald-400">Zero Dependencies</h3>
                <p className="text-sm text-gray-300 text-center">
                  Built entirely on Go's standard library. No external packages required.
                </p>
              </div>
              <div className="flex flex-col items-center space-y-2 p-6 rounded-lg bg-gray-800 transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(0,226,252,0.3)]">
                <div className="relative p-2 rounded-lg">
                  <div className="absolute inset-0 rounded-lg bg-cyan-400/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300 blur-md"></div>
                  <Rocket className="h-8 w-8 text-cyan-400 relative z-10" />
                </div>
                <h3 className="text-xl font-bold text-cyan-400">Standard Compliant</h3>
                <p className="text-sm text-gray-300 text-center">
                  Follows Go idioms and best practices for a familiar development experience.
                </p>
              </div>
              <div className="flex flex-col items-center space-y-2 p-6 rounded-lg bg-gray-800 transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(20,255,180,0.3)]">
                <div className="relative p-2 rounded-lg">
                  <div className="absolute inset-0 rounded-lg bg-teal-400/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300 blur-md"></div>
                  <Zap className="h-8 w-8 text-teal-400 relative z-10" />
                </div>
                <h3 className="text-xl font-bold text-teal-400">Lightweight & Fast</h3>
                <p className="text-sm text-gray-300 text-center">
                  Minimal overhead means blazing fast performance and small binary sizes.
                </p>
              </div>
              <div className="flex flex-col items-center space-y-2 p-6 rounded-lg bg-gray-800 transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(100,255,100,0.3)]">
                <div className="relative p-2 rounded-lg">
                  <div className="absolute inset-0 rounded-lg bg-green-400/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300 blur-md"></div>
                  <Unplug className="h-8 w-8 text-green-400 relative z-10" />
                </div>
                <h3 className="text-xl font-bold text-green-400">WebSocket Support</h3>
                <p className="text-sm text-gray-300 text-center">
                  Seamlessly integrate real-time communication with built-in WebSocket capabilities.
                </p>
                <span className="inline-block bg-yellow-500 text-gray-900 text-xs font-semibold px-2 py-1 rounded-full mt-2">Coming Soon</span>
              </div>
              <div className="flex flex-col items-center space-y-2 p-6 rounded-lg bg-gray-800 transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(255,214,0,0.3)]">
                <div className="relative p-2 rounded-lg">
                  <div className="absolute inset-0 rounded-lg bg-yellow-400/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300 blur-md"></div>
                  <FileJson className="h-8 w-8 text-yellow-400 relative z-10" />
                </div>
                <h3 className="text-xl font-bold text-yellow-400">OpenAPI Integration</h3>
                <p className="text-sm text-gray-300 text-center">
                  Automatically generate OpenAPI specifications for your APIs, enhancing documentation and interoperability.
                </p>
              </div>
              <div className="flex flex-col items-center space-y-2 p-6 rounded-lg bg-gray-800 transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(255,100,100,0.3)]">
                <div className="relative p-2 rounded-lg">
                  <div className="absolute inset-0 rounded-lg bg-pink-400/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300 blur-md"></div>
                  <FileType2 className="h-8 w-8 text-pink-400 relative z-10" />
                </div>
                <h3 className="text-xl font-bold text-pink-400">TypeScript API Generator</h3>
                <p className="text-sm text-gray-300 text-center">
                  Automatically generate TypeScript definitions for your Go APIs, ensuring type safety across your full-stack application.
                </p>
              </div>
            </div>
          </div>
        </section>
        <section id="example" className="w-full py-12 md:py-24 lg:py-32 bg-gray-900 ">
          <div className="container px-4 md:px-6">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-8 bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">Simple to Use</h2>
            <div className="relative max-w-3xl mx-auto">
              <pre className="p-4 bg-gray-800 rounded-lg overflow-x-auto border border-gray-700 shadow-lg transition-all duration-300 ease-in-out hover:bg-gray-800/50 group hover:shadow-[0_0_15px_rgba(20,255,180,0.3)]">
                <code className="text-sm text-gray-300">
                  <button
                    className="absolute top-2 right-2 p-1 rounded-md bg-gray-700 hover:bg-gray-600 transition-colors"
                  >
                    <Copy className="h-4 w-4 text-gray-400" />
                  </button>
                  <div dangerouslySetInnerHTML={{
                    __html: `<pre style="font-family:monospace;color: rgb(197, 200, 198); font-weight: 400; "><span style="color: rgb(34 211 238); font-weight: 400;">package</span> main

<span style="color: rgb(34 211 238); font-weight: 400;">import</span> (
  <span style="color: rgb(74 222 128); font-weight: 400;">"log"</span>
  <span style="color: rgb(74 222 128); font-weight: 400;">"net/http"</span>
  <span style="color: rgb(74 222 128); font-weight: 400;">"github.com/Instantan/web"</span>
)

<span style="color: rgb(197, 200, 198); font-weight: 400;"><span style="color: rgb(34 211 238); font-weight: 400;">func</span> <span class="text-yellow-400" style="font-weight: 400;">main</span><span style="color: rgb(197, 200, 198); font-weight: 400;">()</span></span> {
  w := web.NewWeb()

  w.Info(web.Info{
    Title: <span style="color: rgb(74 222 128); font-weight: 400;">"MyProject"</span>,
    Version: <span style="color: rgb(74 222 128); font-weight: 400;">"0.0.1"</span>,
  })

  w.OpenApi(web.OpenApi{
    DocPath:   <span style="color: rgb(74 222 128); font-weight: 400;">"/api/doc.json"</span>,
    UiPath:    <span style="color: rgb(74 222 128); font-weight: 400;">"/api/doc"</span>,
    UiVariant: <span style="color: rgb(74 222 128); font-weight: 400;">"scalar"</span>,
  })

  w.Api(web.Api{
    Method: http.MethodGet,
    Path:   <span style="color: rgb(74 222 128); font-weight: 400;">"/hello/{name}"</span>,
    Parameter: web.Parameter{
      Path: web.Path{
        <span style="color: rgb(74 222 128); font-weight: 400;">"name"</span>: web.PathParam{
          Description: <span style="color: rgb(74 222 128); font-weight: 400;">"The name to say hello to"</span>,
          Value:       <span style="color: rgb(74 222 128); font-weight: 400;">"world"</span>,
        },
      },
    },
    Responses: web.Responses{
      StatusOK: <span style="color: rgb(74 222 128); font-weight: 400;">"Hello World"</span>,
    },
    Handler: http.HandlerFunc(<span style="color: rgb(197, 200, 198); font-weight: 400;"><span style="color: rgb(34 211 238); font-weight: 400;">func</span><span style="color: rgb(197, 200, 198); font-weight: 400;">(w http.ResponseWriter, r *http.Request)</span></span> {
      w.Write([]<span style="color: rgb(244 114 182); font-weight: 400;">byte</span>(<span style="color: rgb(74 222 128); font-weight: 400;">"Hello "</span> + r.PathValue(<span style="color: rgb(74 222 128); font-weight: 400;">"name"</span>)))
    }),
  })

  log.Println(<span style="color: rgb(74 222 128); font-weight: 400;">"Server listening on :8080"</span>)
  log.Println(<span style="color: rgb(74 222 128); font-weight: 400;">"Visit http://localhost:8080/api/doc to view the documentation"</span>)
  <span style="color: rgb(34 211 238); font-weight: 400;">if</span> err := http.ListenAndServe(<span style="color: rgb(74 222 128); font-weight: 400;">":8080"</span>, w.Server()); err != <span style="color: rgb(244 114 182); font-weight: 400;">nil</span> {
    <span style="color: rgb(244 114 182); font-weight: 400;">panic</span>(err)
  }
}</pre>`
                  }} />
                </code>
              </pre>
            </div>
          </div>
        </section>
        <section className="w-full py-12 md:py-24 lg:py-32 bg-default">
          <div className="container px-4 md:px-6">
            <div className="flex flex-col items-center space-y-4 text-center">
              <div className="space-y-2">
                <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">
                  Ready to build with web?
                </h2>
                <p className="mx-auto max-w-[600px] text-gray-400 md:text-xl">
                  Get started with web today and experience the power of Go's standard library with modern features.
                </p>
              </div>
              <div className="space-x-4">
                <Button className="bg-emerald-500 hover:bg-emerald-600 text-gray-950">Get Started</Button>
                <Button variant="outline" className="border-emerald-500 text-emerald-400 hover:bg-emerald-950">Read the Docs</Button>
              </div>
            </div>
          </div>
        </section>
      </main>
      <footer className="flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6 border-t border-gray-800">
        <p className="text-xs text-gray-400">
          Gopher image "gopher-side_color.svg" by Takuya Ueda (
          <a href="https://twitter.com/tenntenn" className="text-emerald-400 hover:underline">@tenntenn</a>
          ). Licensed under the Creative Commons 3.0 Attributions license.
        </p>
        <div className="sm:ml-auto">
          <a className="text-xs hover:text-emerald-400 transition-colors" href="https://github.com/Instantan/web">
            GitHub
          </a>
        </div>
      </footer>
    </div>
  )
}