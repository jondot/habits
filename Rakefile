

include FileUtils

BINARY = "habits"
ARCHS = [
  ['linux-arm', ''],
  ['linux-amd64',''],
  ['linux-386',''],
  ['darwin-amd64','']
]

task :package do
  ARCHS.each do |arch, prefix|
    `rm -rf pkg/#{arch}`
    `rm -rf build`
    `mkdir build`
    `mkdir -p pkg/#{arch}`
    bin = go_build("build/#{BINARY}", arch)

    cd("pkg/#{arch}") do
      cp "../../#{bin}", BINARY

      tarball = "#{BINARY}-#{arch}.tar.gz"
      `tar -czf #{tarball} *`
      cp tarball, "../"
    end

    `rm -rf pkg/#{arch}`
    `rm -rf build`
  end
end


def go_build(label, arch)
    system({"GOOS" => arch.split('-')[0], 
          "GOARCH" => arch.split('-')[1]}, 
          "gom build -o #{label}-#{arch}")
    "#{label}-#{arch}"
end

