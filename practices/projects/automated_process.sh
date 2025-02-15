#!/bin/bash
# Function to display the current date & time
display_current_time() {
   echo ""
   echo "Script executed at:" $(date '+%d-%m-%Y %H:%M:%S')
   echo ""
}
# Function to check if Dockerfile exists
check_dockerfile() {
   if [[ -f ./Dockerfile ]]; then
      echo " 'Dockerfile' exists in the current working dir:" $(pwd)
      echo ""
      ls -l
      return 0
   else
      echo "Dockerfile does not exist in the current working dir:" $(pwd)
      return 1
   fi
}
# Function to build docker image and run the container
build_and_run_docker() {
   read -p "Do you want to build the docker image ('yes' or 'no')? :" USER_ACTION
   echo "User entered action is: $USER_ACTION"
   sleep 2

   if [[ $USER_ACTION == "yes" || $USER_ACTION == "y" ]]; then
      echo ""
      echo "Building docker image using date with seconds. Starting shortly..."
      echo ""
      image_name="goimage_$(date '+%d-%m-%Y_%H-%M-%S')"
      cont_name="goapp_$(date '+%d-%m-%Y_%H-%M-%S')"
      echo "Generated image name: $image_name"
      echo "Generated container name: $cont_name"
      echo ""
      docker build -t $image_name .
      docker images
      sleep 5
      echo ""
      echo "Deploying Go app in Docker container..."
      docker run --name $cont_name -p 10000:10000 $image_name
      echo ""
      echo "Checking whether the container is started..."
      docker ps -a
      echo ""
      echo "Checking curl output..."
      curl -ivk http://localhost:10000
      echo ""
      remove_container $cont_name
   else
      echo ""
      echo "User selected no option, no further proceedings..."
   fi
}
# Function to remove the Docker container after testing
remove_container() {
   read -p "Hope the testing is done. Do you want to remove the container ('say yes or no')? " CONFIRMATION
   echo "User entered the input as:" $CONFIRMATION
   if [[ $CONFIRMATION == 'yes' || $CONFIRMATION == 'y' ]]; then
      echo "You are good to delete the container.."
      echo ""
      docker ps -a
      echo "Stopping the container.."
      docker stop $1
      docker ps -a
      echo ""
      echo "Deleting the container.."
      docker rm $1
   elif [[ $CONFIRMATION == 'no' || $CONFIRMATION == 'n' ]]; then
      echo "User opted not to remove the container.."
   else
      echo "Please select 'yes or no' as option.."
   fi
}
# Main script execution via function invoking
display_current_time
if check_dockerfile; then
   build_and_run_docker
else
   echo "No further action as Dockerfile is not found."
fi
