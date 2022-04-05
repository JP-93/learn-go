	ar := []int{}
	for i := 0; i <= 4;{
		i++
		if entrar == "sim"{
			fmt.Scan(&entrar)
			ar = append(ar, i)
		}
	}

	fmt.Println(ar)
