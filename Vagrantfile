Vagrant.configure(2) do |config|

  config.vm.define "mysql01" do |mysql01|
    mysql01.vm.box = "consul/mysql"
    mysql01.vm.network "private_network", ip: "10.0.0.14"
  end

  config.vm.define "mysql02" do |mysql02|
    mysql02.vm.box = "lexsys/mysql"
    mysql02.vm.network "private_network", ip: "10.0.0.11"
  end

  config.vm.define "consul" do |consul|
    consul.vm.box = "lexsys/consul"
    consul.vm.network "private_network", ip: "10.0.0.10"
  end

  config.vm.define "lattice" do |lattice|
    system_ip = ENV["LATTICE_SYSTEM_IP"] || "10.0.0.13"
    system_domain = ENV["LATTICE_SYSTEM_DOMAIN"] || "#{system_ip}.xip.io"
    lattice.vm.network "private_network", ip: system_ip

    lattice_tar_version="v0.2.5"
    lattice_tar_url="https://s3-us-west-2.amazonaws.com/lattice/releases/#{lattice_tar_version}/lattice.tgz"

    lattice.vm.box = "lattice/ubuntu-trusty-64"
    lattice.vm.box_version = '0.2.5'

    lattice.vm.provision "shell" do |s|
      populate_lattice_env_file_script = <<-SCRIPT
        mkdir -pv /var/lattice/setup
        echo "CONSUL_SERVER_IP=#{system_ip}" >> /var/lattice/setup/lattice-environment
        echo "SYSTEM_DOMAIN=#{system_domain}" >> /var/lattice/setup/lattice-environment
        echo "LATTICE_CELL_ID=cell-01" >> /var/lattice/setup/lattice-environment
        echo "GARDEN_EXTERNAL_IP=#{system_ip}" >> /var/lattice/setup/lattice-environment
      SCRIPT

      s.inline = populate_lattice_env_file_script
    end

    lattice.vm.provision "shell" do |s|
      s.inline = "cp /var/lattice/setup/lattice-environment /vagrant/.lattice-environment"
    end

    lattice.vm.provision "shell" do |s|
      s.path = "scripts/install-lattice-from-tar"
      s.args = ["collocated", ENV["VAGRANT_LATTICE_TAR_PATH"].to_s, lattice_tar_url]
    end

    lattice.vm.provision "shell" do |s|
      s.inline = "export $(cat /var/lattice/setup/lattice-environment) && echo \"Lattice is now installed and running. You may target it with the Lattice cli via: ltc target $SYSTEM_DOMAIN\""
    end

  end
end