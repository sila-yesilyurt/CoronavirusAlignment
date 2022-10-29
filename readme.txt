Copy the Alignment folder into your "go/src" directory.

Then, write each of the missing functions in the Functions folder. If you need a subroutine, please feel free to place it in the same file.

To test a function, you should "cd" into the Functions directory and run

go test -run ^TestFunctionName$

where TestFunctionName is based on the specific function you are testing.

For example, after writing each appropriate function, you should call each of the following tests.

go test -run ^TestCountSharedKmers$
go test -run ^TestChange$
go test -run ^TestLCSLength$
go test -run ^TestEditDistance$
go test -run ^TestEditDistanceMatrix$
go test -run ^TestLongestCommonSubsequence$
go test -run ^TestGlobalScoreTable$
go test -run ^TestGlobalAlignment$
go test -run ^TestLocalScoreTable$
go test -run ^TestLocalAlignment$

Note: we are providing a function to produce a frequency map in frequency_map.go, since this was covered in the test materials.

To check that you have passed *all* tests, call "go test".

After you have passed all of these tests, navigate into the parent directory (using "cd .."). We will fill in main.go, after which you can call "go build" and then execute the resulting executable file.

The "Data" folder contains some datasets that we will use; chiefly, a few hemoglobin subunit alpha proteins for animals, as well as the SARS-CoV and SARS-CoV-2 whole genomes. The "Output" folder will contain the results of running some code on these genomes.